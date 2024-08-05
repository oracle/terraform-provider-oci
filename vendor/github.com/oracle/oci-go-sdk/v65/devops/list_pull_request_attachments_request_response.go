// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListPullRequestAttachmentsRequest wrapper for the ListPullRequestAttachments operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListPullRequestAttachments.go.html to see an example of how to use ListPullRequestAttachmentsRequest.
type ListPullRequestAttachmentsRequest struct {

	// unique PullRequest identifier
	PullRequestId *string `mandatory:"true" contributesTo:"path" name:"pullRequestId"`

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListPullRequestAttachmentsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order is ascending. If no value is specified timeCreated is default.
	SortBy ListPullRequestAttachmentsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A filter to return only resources that match the entire file name given.
	FileName *string `mandatory:"false" contributesTo:"query" name:"fileName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListPullRequestAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListPullRequestAttachmentsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListPullRequestAttachmentsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListPullRequestAttachmentsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListPullRequestAttachmentsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListPullRequestAttachmentsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListPullRequestAttachmentsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListPullRequestAttachmentsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListPullRequestAttachmentsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListPullRequestAttachmentsResponse wrapper for the ListPullRequestAttachments operation
type ListPullRequestAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of PullRequestAttachmentCollection instances
	PullRequestAttachmentCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListPullRequestAttachmentsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListPullRequestAttachmentsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListPullRequestAttachmentsSortOrderEnum Enum with underlying type: string
type ListPullRequestAttachmentsSortOrderEnum string

// Set of constants representing the allowable values for ListPullRequestAttachmentsSortOrderEnum
const (
	ListPullRequestAttachmentsSortOrderAsc  ListPullRequestAttachmentsSortOrderEnum = "ASC"
	ListPullRequestAttachmentsSortOrderDesc ListPullRequestAttachmentsSortOrderEnum = "DESC"
)

var mappingListPullRequestAttachmentsSortOrderEnum = map[string]ListPullRequestAttachmentsSortOrderEnum{
	"ASC":  ListPullRequestAttachmentsSortOrderAsc,
	"DESC": ListPullRequestAttachmentsSortOrderDesc,
}

var mappingListPullRequestAttachmentsSortOrderEnumLowerCase = map[string]ListPullRequestAttachmentsSortOrderEnum{
	"asc":  ListPullRequestAttachmentsSortOrderAsc,
	"desc": ListPullRequestAttachmentsSortOrderDesc,
}

// GetListPullRequestAttachmentsSortOrderEnumValues Enumerates the set of values for ListPullRequestAttachmentsSortOrderEnum
func GetListPullRequestAttachmentsSortOrderEnumValues() []ListPullRequestAttachmentsSortOrderEnum {
	values := make([]ListPullRequestAttachmentsSortOrderEnum, 0)
	for _, v := range mappingListPullRequestAttachmentsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestAttachmentsSortOrderEnumStringValues Enumerates the set of values in String for ListPullRequestAttachmentsSortOrderEnum
func GetListPullRequestAttachmentsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListPullRequestAttachmentsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestAttachmentsSortOrderEnum(val string) (ListPullRequestAttachmentsSortOrderEnum, bool) {
	enum, ok := mappingListPullRequestAttachmentsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListPullRequestAttachmentsSortByEnum Enum with underlying type: string
type ListPullRequestAttachmentsSortByEnum string

// Set of constants representing the allowable values for ListPullRequestAttachmentsSortByEnum
const (
	ListPullRequestAttachmentsSortByFilename    ListPullRequestAttachmentsSortByEnum = "fileName"
	ListPullRequestAttachmentsSortByTimecreated ListPullRequestAttachmentsSortByEnum = "timeCreated"
	ListPullRequestAttachmentsSortByCreatedby   ListPullRequestAttachmentsSortByEnum = "createdBy"
)

var mappingListPullRequestAttachmentsSortByEnum = map[string]ListPullRequestAttachmentsSortByEnum{
	"fileName":    ListPullRequestAttachmentsSortByFilename,
	"timeCreated": ListPullRequestAttachmentsSortByTimecreated,
	"createdBy":   ListPullRequestAttachmentsSortByCreatedby,
}

var mappingListPullRequestAttachmentsSortByEnumLowerCase = map[string]ListPullRequestAttachmentsSortByEnum{
	"filename":    ListPullRequestAttachmentsSortByFilename,
	"timecreated": ListPullRequestAttachmentsSortByTimecreated,
	"createdby":   ListPullRequestAttachmentsSortByCreatedby,
}

// GetListPullRequestAttachmentsSortByEnumValues Enumerates the set of values for ListPullRequestAttachmentsSortByEnum
func GetListPullRequestAttachmentsSortByEnumValues() []ListPullRequestAttachmentsSortByEnum {
	values := make([]ListPullRequestAttachmentsSortByEnum, 0)
	for _, v := range mappingListPullRequestAttachmentsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListPullRequestAttachmentsSortByEnumStringValues Enumerates the set of values in String for ListPullRequestAttachmentsSortByEnum
func GetListPullRequestAttachmentsSortByEnumStringValues() []string {
	return []string{
		"fileName",
		"timeCreated",
		"createdBy",
	}
}

// GetMappingListPullRequestAttachmentsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListPullRequestAttachmentsSortByEnum(val string) (ListPullRequestAttachmentsSortByEnum, bool) {
	enum, ok := mappingListPullRequestAttachmentsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
