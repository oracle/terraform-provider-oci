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

// ListOdaPrivateEndpointAttachmentsRequest wrapper for the ListOdaPrivateEndpointAttachments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/oda/ListOdaPrivateEndpointAttachments.go.html to see an example of how to use ListOdaPrivateEndpointAttachmentsRequest.
type ListOdaPrivateEndpointAttachmentsRequest struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of ODA Private Endpoint.
	OdaPrivateEndpointId *string `mandatory:"true" contributesTo:"query" name:"odaPrivateEndpointId"`

	// List the ODA Private Endpoint Attachments that belong to this compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// List only the ODA Private Endpoint Attachments that are in this lifecycle state.
	LifecycleState OdaPrivateEndpointAttachmentLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListOdaPrivateEndpointAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `TIMECREATED`.
	// The default sort order for `TIMECREATED` is descending, and the default sort order for `DISPLAYNAME` is ascending.
	SortBy ListOdaPrivateEndpointAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListOdaPrivateEndpointAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListOdaPrivateEndpointAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListOdaPrivateEndpointAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListOdaPrivateEndpointAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListOdaPrivateEndpointAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaPrivateEndpointAttachmentLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetOdaPrivateEndpointAttachmentLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaPrivateEndpointAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListOdaPrivateEndpointAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListOdaPrivateEndpointAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListOdaPrivateEndpointAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListOdaPrivateEndpointAttachmentsResponse wrapper for the ListOdaPrivateEndpointAttachments operation
type ListOdaPrivateEndpointAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of OdaPrivateEndpointAttachmentCollection instances
	OdaPrivateEndpointAttachmentCollection `presentIn:"body"`

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

func (response ListOdaPrivateEndpointAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListOdaPrivateEndpointAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListOdaPrivateEndpointAttachmentsSortOrderEnum Enum with underlying type: string
type ListOdaPrivateEndpointAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListOdaPrivateEndpointAttachmentsSortOrderEnum
const (
	ListOdaPrivateEndpointAttachmentsSortOrderAsc  ListOdaPrivateEndpointAttachmentsSortOrderEnum = "ASC"
	ListOdaPrivateEndpointAttachmentsSortOrderDesc ListOdaPrivateEndpointAttachmentsSortOrderEnum = "DESC"
)

var mappingListOdaPrivateEndpointAttachmentsSortOrderEnum = map[string]ListOdaPrivateEndpointAttachmentsSortOrderEnum{
	"ASC":  ListOdaPrivateEndpointAttachmentsSortOrderAsc,
	"DESC": ListOdaPrivateEndpointAttachmentsSortOrderDesc,
}

var mappingListOdaPrivateEndpointAttachmentsSortOrderEnumLowerCase = map[string]ListOdaPrivateEndpointAttachmentsSortOrderEnum{
	"asc":  ListOdaPrivateEndpointAttachmentsSortOrderAsc,
	"desc": ListOdaPrivateEndpointAttachmentsSortOrderDesc,
}

// GetListOdaPrivateEndpointAttachmentsSortOrderEnumValues Enumerates the set of values for ListOdaPrivateEndpointAttachmentsSortOrderEnum
func GetListOdaPrivateEndpointAttachmentsSortOrderEnumValues() []ListOdaPrivateEndpointAttachmentsSortOrderEnum {
	values := make([]ListOdaPrivateEndpointAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListOdaPrivateEndpointAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaPrivateEndpointAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListOdaPrivateEndpointAttachmentsSortOrderEnum
func GetListOdaPrivateEndpointAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListOdaPrivateEndpointAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaPrivateEndpointAttachmentsSortOrderEnum(val string) (ListOdaPrivateEndpointAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListOdaPrivateEndpointAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListOdaPrivateEndpointAttachmentsSortByEnum Enum with underlying type: string
type ListOdaPrivateEndpointAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListOdaPrivateEndpointAttachmentsSortByEnum
const (
	ListOdaPrivateEndpointAttachmentsSortByTimecreated ListOdaPrivateEndpointAttachmentsSortByEnum = "TIMECREATED"
	ListOdaPrivateEndpointAttachmentsSortByDisplayname ListOdaPrivateEndpointAttachmentsSortByEnum = "DISPLAYNAME"
)

var mappingListOdaPrivateEndpointAttachmentsSortByEnum = map[string]ListOdaPrivateEndpointAttachmentsSortByEnum{
	"TIMECREATED": ListOdaPrivateEndpointAttachmentsSortByTimecreated,
	"DISPLAYNAME": ListOdaPrivateEndpointAttachmentsSortByDisplayname,
}

var mappingListOdaPrivateEndpointAttachmentsSortByEnumLowerCase = map[string]ListOdaPrivateEndpointAttachmentsSortByEnum{
	"timecreated": ListOdaPrivateEndpointAttachmentsSortByTimecreated,
	"displayname": ListOdaPrivateEndpointAttachmentsSortByDisplayname,
}

// GetListOdaPrivateEndpointAttachmentsSortByEnumValues Enumerates the set of values for ListOdaPrivateEndpointAttachmentsSortByEnum
func GetListOdaPrivateEndpointAttachmentsSortByEnumValues() []ListOdaPrivateEndpointAttachmentsSortByEnum {
	values := make([]ListOdaPrivateEndpointAttachmentsSortByEnum, 0)
	for _, v := range mappingListOdaPrivateEndpointAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListOdaPrivateEndpointAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListOdaPrivateEndpointAttachmentsSortByEnum
func GetListOdaPrivateEndpointAttachmentsSortByEnumStringValues() []string {
	return []string{
		"TIMECREATED",
		"DISPLAYNAME",
	}
}

// GetMappingListOdaPrivateEndpointAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListOdaPrivateEndpointAttachmentsSortByEnum(val string) (ListOdaPrivateEndpointAttachmentsSortByEnum, bool) {
	enum, ok := mappingListOdaPrivateEndpointAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
