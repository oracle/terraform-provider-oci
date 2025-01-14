// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListSkillKnowledgeGroupAssociationsRequest wrapper for the ListSkillKnowledgeGroupAssociations operation
type ListSkillKnowledgeGroupAssociationsRequest struct {

	// Unique Digital Assistant instance identifier.
	OdaInstanceId *string `mandatory:"true" contributesTo:"path" name:"odaInstanceId"`

	// Unique Skill identifier.
	SkillId *string `mandatory:"true" contributesTo:"path" name:"skillId"`

	// The maximum number of items to return per page.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page at which to start retrieving results.
	// You get this value from the `opc-next-page` header in a previous list request.
	// To retireve the first page, omit this query parameter.
	// Example: `MToxMA==`
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Sort the results in this order, use either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListSkillKnowledgeGroupAssociationsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Sort on this field. You can specify one sort order only. The default sort field is `timeCreated`.
	// The default sort order for `timeCreated` and `timeUpdated` is descending
	SortBy ListSkillKnowledgeGroupAssociationsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing. This value is included in the opc-request-id response header.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSkillKnowledgeGroupAssociationsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSkillKnowledgeGroupAssociationsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSkillKnowledgeGroupAssociationsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSkillKnowledgeGroupAssociationsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSkillKnowledgeGroupAssociationsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSkillKnowledgeGroupAssociationsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListSkillKnowledgeGroupAssociationsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSkillKnowledgeGroupAssociationsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListSkillKnowledgeGroupAssociationsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSkillKnowledgeGroupAssociationsResponse wrapper for the ListSkillKnowledgeGroupAssociations operation
type ListSkillKnowledgeGroupAssociationsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SkillKnowledgeGroupAssociationCollection instances
	SkillKnowledgeGroupAssociationCollection `presentIn:"body"`

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

func (response ListSkillKnowledgeGroupAssociationsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSkillKnowledgeGroupAssociationsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSkillKnowledgeGroupAssociationsSortOrderEnum Enum with underlying type: string
type ListSkillKnowledgeGroupAssociationsSortOrderEnum string

// Set of constants representing the allowable values for ListSkillKnowledgeGroupAssociationsSortOrderEnum
const (
	ListSkillKnowledgeGroupAssociationsSortOrderAsc  ListSkillKnowledgeGroupAssociationsSortOrderEnum = "ASC"
	ListSkillKnowledgeGroupAssociationsSortOrderDesc ListSkillKnowledgeGroupAssociationsSortOrderEnum = "DESC"
)

var mappingListSkillKnowledgeGroupAssociationsSortOrderEnum = map[string]ListSkillKnowledgeGroupAssociationsSortOrderEnum{
	"ASC":  ListSkillKnowledgeGroupAssociationsSortOrderAsc,
	"DESC": ListSkillKnowledgeGroupAssociationsSortOrderDesc,
}

var mappingListSkillKnowledgeGroupAssociationsSortOrderEnumLowerCase = map[string]ListSkillKnowledgeGroupAssociationsSortOrderEnum{
	"asc":  ListSkillKnowledgeGroupAssociationsSortOrderAsc,
	"desc": ListSkillKnowledgeGroupAssociationsSortOrderDesc,
}

// GetListSkillKnowledgeGroupAssociationsSortOrderEnumValues Enumerates the set of values for ListSkillKnowledgeGroupAssociationsSortOrderEnum
func GetListSkillKnowledgeGroupAssociationsSortOrderEnumValues() []ListSkillKnowledgeGroupAssociationsSortOrderEnum {
	values := make([]ListSkillKnowledgeGroupAssociationsSortOrderEnum, 0)
	for _, v := range mappingListSkillKnowledgeGroupAssociationsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillKnowledgeGroupAssociationsSortOrderEnumStringValues Enumerates the set of values in String for ListSkillKnowledgeGroupAssociationsSortOrderEnum
func GetListSkillKnowledgeGroupAssociationsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListSkillKnowledgeGroupAssociationsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillKnowledgeGroupAssociationsSortOrderEnum(val string) (ListSkillKnowledgeGroupAssociationsSortOrderEnum, bool) {
	enum, ok := mappingListSkillKnowledgeGroupAssociationsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSkillKnowledgeGroupAssociationsSortByEnum Enum with underlying type: string
type ListSkillKnowledgeGroupAssociationsSortByEnum string

// Set of constants representing the allowable values for ListSkillKnowledgeGroupAssociationsSortByEnum
const (
	ListSkillKnowledgeGroupAssociationsSortByTimecreated ListSkillKnowledgeGroupAssociationsSortByEnum = "timeCreated"
	ListSkillKnowledgeGroupAssociationsSortByTimeupdated ListSkillKnowledgeGroupAssociationsSortByEnum = "timeUpdated"
)

var mappingListSkillKnowledgeGroupAssociationsSortByEnum = map[string]ListSkillKnowledgeGroupAssociationsSortByEnum{
	"timeCreated": ListSkillKnowledgeGroupAssociationsSortByTimecreated,
	"timeUpdated": ListSkillKnowledgeGroupAssociationsSortByTimeupdated,
}

var mappingListSkillKnowledgeGroupAssociationsSortByEnumLowerCase = map[string]ListSkillKnowledgeGroupAssociationsSortByEnum{
	"timecreated": ListSkillKnowledgeGroupAssociationsSortByTimecreated,
	"timeupdated": ListSkillKnowledgeGroupAssociationsSortByTimeupdated,
}

// GetListSkillKnowledgeGroupAssociationsSortByEnumValues Enumerates the set of values for ListSkillKnowledgeGroupAssociationsSortByEnum
func GetListSkillKnowledgeGroupAssociationsSortByEnumValues() []ListSkillKnowledgeGroupAssociationsSortByEnum {
	values := make([]ListSkillKnowledgeGroupAssociationsSortByEnum, 0)
	for _, v := range mappingListSkillKnowledgeGroupAssociationsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListSkillKnowledgeGroupAssociationsSortByEnumStringValues Enumerates the set of values in String for ListSkillKnowledgeGroupAssociationsSortByEnum
func GetListSkillKnowledgeGroupAssociationsSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"timeUpdated",
	}
}

// GetMappingListSkillKnowledgeGroupAssociationsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSkillKnowledgeGroupAssociationsSortByEnum(val string) (ListSkillKnowledgeGroupAssociationsSortByEnum, bool) {
	enum, ok := mappingListSkillKnowledgeGroupAssociationsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
