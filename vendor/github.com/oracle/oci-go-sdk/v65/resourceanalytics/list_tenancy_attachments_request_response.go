// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListTenancyAttachmentsRequest wrapper for the ListTenancyAttachments operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/resourceanalytics/ListTenancyAttachments.go.html to see an example of how to use ListTenancyAttachmentsRequest.
type ListTenancyAttachmentsRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a ResourceAnalyticsInstance.
	ResourceAnalyticsInstanceId *string `mandatory:"false" contributesTo:"query" name:"resourceAnalyticsInstanceId"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState TenancyAttachmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the TenancyAttachment.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// For list pagination. The maximum number of results per page, or items to return in a
	// paginated "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the opc-next-page response header from the previous
	// "List" call. For important details about how pagination works, see
	// List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListTenancyAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. You can provide only one sort order. Default order for `TIME_CREATED`
	// is descending. Default order for `DESCRIPTION` is ascending.
	SortBy ListTenancyAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	// The only valid characters for request IDs are letters, numbers,
	// underscore, and dash.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListTenancyAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListTenancyAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListTenancyAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListTenancyAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListTenancyAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTenancyAttachmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetTenancyAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTenancyAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListTenancyAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListTenancyAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListTenancyAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListTenancyAttachmentsResponse wrapper for the ListTenancyAttachments operation
type ListTenancyAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TenancyAttachmentCollection instances
	TenancyAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. For
	// important details about how pagination works, see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListTenancyAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListTenancyAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListTenancyAttachmentsSortOrderEnum Enum with underlying type: string
type ListTenancyAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListTenancyAttachmentsSortOrderEnum
const (
	ListTenancyAttachmentsSortOrderAsc  ListTenancyAttachmentsSortOrderEnum = "ASC"
	ListTenancyAttachmentsSortOrderDesc ListTenancyAttachmentsSortOrderEnum = "DESC"
)

var mappingListTenancyAttachmentsSortOrderEnum = map[string]ListTenancyAttachmentsSortOrderEnum{
	"ASC":  ListTenancyAttachmentsSortOrderAsc,
	"DESC": ListTenancyAttachmentsSortOrderDesc,
}

var mappingListTenancyAttachmentsSortOrderEnumLowerCase = map[string]ListTenancyAttachmentsSortOrderEnum{
	"asc":  ListTenancyAttachmentsSortOrderAsc,
	"desc": ListTenancyAttachmentsSortOrderDesc,
}

// GetListTenancyAttachmentsSortOrderEnumValues Enumerates the set of values for ListTenancyAttachmentsSortOrderEnum
func GetListTenancyAttachmentsSortOrderEnumValues() []ListTenancyAttachmentsSortOrderEnum {
	values := make([]ListTenancyAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListTenancyAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListTenancyAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListTenancyAttachmentsSortOrderEnum
func GetListTenancyAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListTenancyAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTenancyAttachmentsSortOrderEnum(val string) (ListTenancyAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListTenancyAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListTenancyAttachmentsSortByEnum Enum with underlying type: string
type ListTenancyAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListTenancyAttachmentsSortByEnum
const (
	ListTenancyAttachmentsSortByTimeCreated ListTenancyAttachmentsSortByEnum = "TIME_CREATED"
	ListTenancyAttachmentsSortByDescription ListTenancyAttachmentsSortByEnum = "DESCRIPTION"
)

var mappingListTenancyAttachmentsSortByEnum = map[string]ListTenancyAttachmentsSortByEnum{
	"TIME_CREATED": ListTenancyAttachmentsSortByTimeCreated,
	"DESCRIPTION":  ListTenancyAttachmentsSortByDescription,
}

var mappingListTenancyAttachmentsSortByEnumLowerCase = map[string]ListTenancyAttachmentsSortByEnum{
	"time_created": ListTenancyAttachmentsSortByTimeCreated,
	"description":  ListTenancyAttachmentsSortByDescription,
}

// GetListTenancyAttachmentsSortByEnumValues Enumerates the set of values for ListTenancyAttachmentsSortByEnum
func GetListTenancyAttachmentsSortByEnumValues() []ListTenancyAttachmentsSortByEnum {
	values := make([]ListTenancyAttachmentsSortByEnum, 0)
	for _, v := range mappingListTenancyAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListTenancyAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListTenancyAttachmentsSortByEnum
func GetListTenancyAttachmentsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"DESCRIPTION",
	}
}

// GetMappingListTenancyAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListTenancyAttachmentsSortByEnum(val string) (ListTenancyAttachmentsSortByEnum, bool) {
	enum, ok := mappingListTenancyAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
