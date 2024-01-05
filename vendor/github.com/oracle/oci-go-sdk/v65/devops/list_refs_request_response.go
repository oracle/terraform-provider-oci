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

// ListRefsRequest wrapper for the ListRefs operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListRefs.go.html to see an example of how to use ListRefsRequest.
type ListRefsRequest struct {

	// Unique repository identifier.
	RepositoryId *string `mandatory:"true" contributesTo:"path" name:"repositoryId"`

	// Reference type to distinguish between branch and tag. If it is not specified, all references are returned.
	RefType ListRefsRefTypeEnum `mandatory:"false" contributesTo:"query" name:"refType" omitEmpty:"true"`

	// Commit ID in a repository.
	CommitId *string `mandatory:"false" contributesTo:"query" name:"commitId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the given reference name.
	RefName *string `mandatory:"false" contributesTo:"query" name:"refName"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRefsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for reference name is ascending. Default order for reference type is ascending. If no value is specified reference name is default.
	SortBy ListRefsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRefsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRefsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRefsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRefsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListRefsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListRefsRefTypeEnum(string(request.RefType)); !ok && request.RefType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RefType: %s. Supported values are: %s.", request.RefType, strings.Join(GetListRefsRefTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRefsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListRefsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListRefsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListRefsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListRefsResponse wrapper for the ListRefs operation
type ListRefsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryRefCollection instances
	RepositoryRefCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRefsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRefsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRefsRefTypeEnum Enum with underlying type: string
type ListRefsRefTypeEnum string

// Set of constants representing the allowable values for ListRefsRefTypeEnum
const (
	ListRefsRefTypeBranch ListRefsRefTypeEnum = "BRANCH"
	ListRefsRefTypeTag    ListRefsRefTypeEnum = "TAG"
)

var mappingListRefsRefTypeEnum = map[string]ListRefsRefTypeEnum{
	"BRANCH": ListRefsRefTypeBranch,
	"TAG":    ListRefsRefTypeTag,
}

var mappingListRefsRefTypeEnumLowerCase = map[string]ListRefsRefTypeEnum{
	"branch": ListRefsRefTypeBranch,
	"tag":    ListRefsRefTypeTag,
}

// GetListRefsRefTypeEnumValues Enumerates the set of values for ListRefsRefTypeEnum
func GetListRefsRefTypeEnumValues() []ListRefsRefTypeEnum {
	values := make([]ListRefsRefTypeEnum, 0)
	for _, v := range mappingListRefsRefTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListRefsRefTypeEnumStringValues Enumerates the set of values in String for ListRefsRefTypeEnum
func GetListRefsRefTypeEnumStringValues() []string {
	return []string{
		"BRANCH",
		"TAG",
	}
}

// GetMappingListRefsRefTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRefsRefTypeEnum(val string) (ListRefsRefTypeEnum, bool) {
	enum, ok := mappingListRefsRefTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRefsSortOrderEnum Enum with underlying type: string
type ListRefsSortOrderEnum string

// Set of constants representing the allowable values for ListRefsSortOrderEnum
const (
	ListRefsSortOrderAsc  ListRefsSortOrderEnum = "ASC"
	ListRefsSortOrderDesc ListRefsSortOrderEnum = "DESC"
)

var mappingListRefsSortOrderEnum = map[string]ListRefsSortOrderEnum{
	"ASC":  ListRefsSortOrderAsc,
	"DESC": ListRefsSortOrderDesc,
}

var mappingListRefsSortOrderEnumLowerCase = map[string]ListRefsSortOrderEnum{
	"asc":  ListRefsSortOrderAsc,
	"desc": ListRefsSortOrderDesc,
}

// GetListRefsSortOrderEnumValues Enumerates the set of values for ListRefsSortOrderEnum
func GetListRefsSortOrderEnumValues() []ListRefsSortOrderEnum {
	values := make([]ListRefsSortOrderEnum, 0)
	for _, v := range mappingListRefsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListRefsSortOrderEnumStringValues Enumerates the set of values in String for ListRefsSortOrderEnum
func GetListRefsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListRefsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRefsSortOrderEnum(val string) (ListRefsSortOrderEnum, bool) {
	enum, ok := mappingListRefsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListRefsSortByEnum Enum with underlying type: string
type ListRefsSortByEnum string

// Set of constants representing the allowable values for ListRefsSortByEnum
const (
	ListRefsSortByReftype ListRefsSortByEnum = "refType"
	ListRefsSortByRefname ListRefsSortByEnum = "refName"
)

var mappingListRefsSortByEnum = map[string]ListRefsSortByEnum{
	"refType": ListRefsSortByReftype,
	"refName": ListRefsSortByRefname,
}

var mappingListRefsSortByEnumLowerCase = map[string]ListRefsSortByEnum{
	"reftype": ListRefsSortByReftype,
	"refname": ListRefsSortByRefname,
}

// GetListRefsSortByEnumValues Enumerates the set of values for ListRefsSortByEnum
func GetListRefsSortByEnumValues() []ListRefsSortByEnum {
	values := make([]ListRefsSortByEnum, 0)
	for _, v := range mappingListRefsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListRefsSortByEnumStringValues Enumerates the set of values in String for ListRefsSortByEnum
func GetListRefsSortByEnumStringValues() []string {
	return []string{
		"refType",
		"refName",
	}
}

// GetMappingListRefsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListRefsSortByEnum(val string) (ListRefsSortByEnum, bool) {
	enum, ok := mappingListRefsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
