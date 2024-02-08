// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListOdaInstanceAttachmentsRequest wrapper for the ListOdaInstanceAttachments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListOdaInstanceAttachments.go.html to see an example of how to use ListOdaInstanceAttachmentsRequest.
type ListOdaInstanceAttachmentsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Whether to send attachment owner info during get/list call.
	IncludeOwnerMetadata *bool `mandatory:"false" contributesTo:"query" name:"includeOwnerMetadata"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// List only the ODA instance attachments that are in this lifecycle state.
	LifecycleState ListOdaInstanceAttachmentsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListOdaInstanceAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `TIMECREATED`.
	// The default sort order for `TIMECREATED` is descending.
	SortBy ListOdaInstanceAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOdaInstanceAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOdaInstanceAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOdaInstanceAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOdaInstanceAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOdaInstanceAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListOdaInstanceAttachmentsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListOdaInstanceAttachmentsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaInstanceAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOdaInstanceAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaInstanceAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOdaInstanceAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOdaInstanceAttachmentsResponse wrapper for the ListOdaInstanceAttachments operation
type ListOdaInstanceAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OdaInstanceAttachmentCollection instances
	OdaInstanceAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// When you are paging through a list, if this header appears in the response,
	// then there might be additional items still to get. Include this value as the
	// `page` query parameter for the subsequent GET request.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// The total number of results that match the query.
	OpcTotalItems *int `presentIn:"header" name:"opc-total-items"`
}

func (response ListOdaInstanceAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOdaInstanceAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOdaInstanceAttachmentsLifecycleStateEnum Enum with underlying type: string
type ListOdaInstanceAttachmentsLifecycleStateEnum string

// Set of constants representing the allowable values for ListOdaInstanceAttachmentsLifecycleStateEnum
const (
	ListOdaInstanceAttachmentsLifecycleStateAttaching ListOdaInstanceAttachmentsLifecycleStateEnum = "ATTACHING"
	ListOdaInstanceAttachmentsLifecycleStateActive    ListOdaInstanceAttachmentsLifecycleStateEnum = "ACTIVE"
	ListOdaInstanceAttachmentsLifecycleStateDetaching ListOdaInstanceAttachmentsLifecycleStateEnum = "DETACHING"
	ListOdaInstanceAttachmentsLifecycleStateInactive  ListOdaInstanceAttachmentsLifecycleStateEnum = "INACTIVE"
	ListOdaInstanceAttachmentsLifecycleStateFailed    ListOdaInstanceAttachmentsLifecycleStateEnum = "FAILED"
)

var mappingListOdaInstanceAttachmentsLifecycleStateEnum = map[string]ListOdaInstanceAttachmentsLifecycleStateEnum{
	"ATTACHING": ListOdaInstanceAttachmentsLifecycleStateAttaching,
	"ACTIVE":    ListOdaInstanceAttachmentsLifecycleStateActive,
	"DETACHING": ListOdaInstanceAttachmentsLifecycleStateDetaching,
	"INACTIVE":  ListOdaInstanceAttachmentsLifecycleStateInactive,
	"FAILED":    ListOdaInstanceAttachmentsLifecycleStateFailed,
}

var mappingListOdaInstanceAttachmentsLifecycleStateEnumLowerCase = map[string]ListOdaInstanceAttachmentsLifecycleStateEnum{
	"attaching": ListOdaInstanceAttachmentsLifecycleStateAttaching,
	"active":    ListOdaInstanceAttachmentsLifecycleStateActive,
	"detaching": ListOdaInstanceAttachmentsLifecycleStateDetaching,
	"inactive":  ListOdaInstanceAttachmentsLifecycleStateInactive,
	"failed":    ListOdaInstanceAttachmentsLifecycleStateFailed,
}

// GetListOdaInstanceAttachmentsLifecycleStateEnumValues Enumerates the set of values for ListOdaInstanceAttachmentsLifecycleStateEnum
func GetListOdaInstanceAttachmentsLifecycleStateEnumValues() []ListOdaInstanceAttachmentsLifecycleStateEnum {
	values := make([]ListOdaInstanceAttachmentsLifecycleStateEnum, 0)
	for _, v := range mappingListOdaInstanceAttachmentsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaInstanceAttachmentsLifecycleStateEnumStringValues Enumerates the set of values in String for ListOdaInstanceAttachmentsLifecycleStateEnum
func GetListOdaInstanceAttachmentsLifecycleStateEnumStringValues() []string {
	return []string{
		"ATTACHING",
		"ACTIVE",
		"DETACHING",
		"INACTIVE",
		"FAILED",
	}
}

// GetMappingListOdaInstanceAttachmentsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaInstanceAttachmentsLifecycleStateEnum(val string) (ListOdaInstanceAttachmentsLifecycleStateEnum, bool) {
	enum, ok := mappingListOdaInstanceAttachmentsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOdaInstanceAttachmentsSortOrderEnum Enum with underlying type: string
type ListOdaInstanceAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListOdaInstanceAttachmentsSortOrderEnum
const (
	ListOdaInstanceAttachmentsSortOrderAsc  ListOdaInstanceAttachmentsSortOrderEnum = "ASC"
	ListOdaInstanceAttachmentsSortOrderDesc ListOdaInstanceAttachmentsSortOrderEnum = "DESC"
)

var mappingListOdaInstanceAttachmentsSortOrderEnum = map[string]ListOdaInstanceAttachmentsSortOrderEnum{
	"ASC":  ListOdaInstanceAttachmentsSortOrderAsc,
	"DESC": ListOdaInstanceAttachmentsSortOrderDesc,
}

var mappingListOdaInstanceAttachmentsSortOrderEnumLowerCase = map[string]ListOdaInstanceAttachmentsSortOrderEnum{
	"asc":  ListOdaInstanceAttachmentsSortOrderAsc,
	"desc": ListOdaInstanceAttachmentsSortOrderDesc,
}

// GetListOdaInstanceAttachmentsSortOrderEnumValues Enumerates the set of values for ListOdaInstanceAttachmentsSortOrderEnum
func GetListOdaInstanceAttachmentsSortOrderEnumValues() []ListOdaInstanceAttachmentsSortOrderEnum {
	values := make([]ListOdaInstanceAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListOdaInstanceAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaInstanceAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListOdaInstanceAttachmentsSortOrderEnum
func GetListOdaInstanceAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOdaInstanceAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaInstanceAttachmentsSortOrderEnum(val string) (ListOdaInstanceAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListOdaInstanceAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOdaInstanceAttachmentsSortByEnum Enum with underlying type: string
type ListOdaInstanceAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListOdaInstanceAttachmentsSortByEnum
const (
	ListOdaInstanceAttachmentsSortByTimecreated ListOdaInstanceAttachmentsSortByEnum = "TIMECREATED"
)

var mappingListOdaInstanceAttachmentsSortByEnum = map[string]ListOdaInstanceAttachmentsSortByEnum{
	"TIMECREATED": ListOdaInstanceAttachmentsSortByTimecreated,
}

var mappingListOdaInstanceAttachmentsSortByEnumLowerCase = map[string]ListOdaInstanceAttachmentsSortByEnum{
	"timecreated": ListOdaInstanceAttachmentsSortByTimecreated,
}

// GetListOdaInstanceAttachmentsSortByEnumValues Enumerates the set of values for ListOdaInstanceAttachmentsSortByEnum
func GetListOdaInstanceAttachmentsSortByEnumValues() []ListOdaInstanceAttachmentsSortByEnum {
	values := make([]ListOdaInstanceAttachmentsSortByEnum, 0)
	for _, v := range mappingListOdaInstanceAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaInstanceAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListOdaInstanceAttachmentsSortByEnum
func GetListOdaInstanceAttachmentsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
	}
}

// GetMappingListOdaInstanceAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaInstanceAttachmentsSortByEnum(val string) (ListOdaInstanceAttachmentsSortByEnum, bool) {
	enum, ok := mappingListOdaInstanceAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
