// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListExcludedObjectsRequest wrapper for the ListExcludedObjects operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/ListExcludedObjects.go.html to see an example of how to use ListExcludedObjectsRequest.
type ListExcludedObjectsRequest struct {

	// The OCID of the job
	JobId *string `mandatory:"true" contributesTo:"path" name:"jobId"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListExcludedObjectsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for reasonCategory is ascending.
	// If no value is specified reasonCategory is default.
	SortBy ListExcludedObjectsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Excluded object type.
	Type *string `mandatory:"false" contributesTo:"query" name:"type"`

	// Excluded object owner
	Owner *string `mandatory:"false" contributesTo:"query" name:"owner"`

	// Excluded object name
	Object *string `mandatory:"false" contributesTo:"query" name:"object"`

	// Excluded object owner which contains provided value.
	OwnerContains *string `mandatory:"false" contributesTo:"query" name:"ownerContains"`

	// Excluded object name which contains provided value.
	ObjectContains *string `mandatory:"false" contributesTo:"query" name:"objectContains"`

	// Reason category for the excluded object
	ReasonCategory ListExcludedObjectsReasonCategoryEnum `mandatory:"false" contributesTo:"query" name:"reasonCategory" omitEmpty:"true"`

	// Exclude object rule that matches the excluded object, if applicable.
	SourceRule *string `mandatory:"false" contributesTo:"query" name:"sourceRule"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListExcludedObjectsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListExcludedObjectsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListExcludedObjectsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListExcludedObjectsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListExcludedObjectsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListExcludedObjectsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListExcludedObjectsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExcludedObjectsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListExcludedObjectsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListExcludedObjectsReasonCategoryEnum(string(request.ReasonCategory)); !ok && request.ReasonCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReasonCategory: %s. Supported values are: %s.", request.ReasonCategory, strings.Join(GetListExcludedObjectsReasonCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListExcludedObjectsResponse wrapper for the ListExcludedObjects operation
type ListExcludedObjectsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of ExcludedObjectSummaryCollection instances
	ExcludedObjectSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListExcludedObjectsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListExcludedObjectsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListExcludedObjectsSortOrderEnum Enum with underlying type: string
type ListExcludedObjectsSortOrderEnum string

// Set of constants representing the allowable values for ListExcludedObjectsSortOrderEnum
const (
	ListExcludedObjectsSortOrderAsc  ListExcludedObjectsSortOrderEnum = "ASC"
	ListExcludedObjectsSortOrderDesc ListExcludedObjectsSortOrderEnum = "DESC"
)

var mappingListExcludedObjectsSortOrderEnum = map[string]ListExcludedObjectsSortOrderEnum{
	"ASC":  ListExcludedObjectsSortOrderAsc,
	"DESC": ListExcludedObjectsSortOrderDesc,
}

var mappingListExcludedObjectsSortOrderEnumLowerCase = map[string]ListExcludedObjectsSortOrderEnum{
	"asc":  ListExcludedObjectsSortOrderAsc,
	"desc": ListExcludedObjectsSortOrderDesc,
}

// GetListExcludedObjectsSortOrderEnumValues Enumerates the set of values for ListExcludedObjectsSortOrderEnum
func GetListExcludedObjectsSortOrderEnumValues() []ListExcludedObjectsSortOrderEnum {
	values := make([]ListExcludedObjectsSortOrderEnum, 0)
	for _, v := range mappingListExcludedObjectsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListExcludedObjectsSortOrderEnumStringValues Enumerates the set of values in String for ListExcludedObjectsSortOrderEnum
func GetListExcludedObjectsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListExcludedObjectsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExcludedObjectsSortOrderEnum(val string) (ListExcludedObjectsSortOrderEnum, bool) {
	enum, ok := mappingListExcludedObjectsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExcludedObjectsSortByEnum Enum with underlying type: string
type ListExcludedObjectsSortByEnum string

// Set of constants representing the allowable values for ListExcludedObjectsSortByEnum
const (
	ListExcludedObjectsSortByType           ListExcludedObjectsSortByEnum = "type"
	ListExcludedObjectsSortByReasoncategory ListExcludedObjectsSortByEnum = "reasonCategory"
)

var mappingListExcludedObjectsSortByEnum = map[string]ListExcludedObjectsSortByEnum{
	"type":           ListExcludedObjectsSortByType,
	"reasonCategory": ListExcludedObjectsSortByReasoncategory,
}

var mappingListExcludedObjectsSortByEnumLowerCase = map[string]ListExcludedObjectsSortByEnum{
	"type":           ListExcludedObjectsSortByType,
	"reasoncategory": ListExcludedObjectsSortByReasoncategory,
}

// GetListExcludedObjectsSortByEnumValues Enumerates the set of values for ListExcludedObjectsSortByEnum
func GetListExcludedObjectsSortByEnumValues() []ListExcludedObjectsSortByEnum {
	values := make([]ListExcludedObjectsSortByEnum, 0)
	for _, v := range mappingListExcludedObjectsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListExcludedObjectsSortByEnumStringValues Enumerates the set of values in String for ListExcludedObjectsSortByEnum
func GetListExcludedObjectsSortByEnumStringValues() []string {
	return []string{
		"type",
		"reasonCategory",
	}
}

// GetMappingListExcludedObjectsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExcludedObjectsSortByEnum(val string) (ListExcludedObjectsSortByEnum, bool) {
	enum, ok := mappingListExcludedObjectsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListExcludedObjectsReasonCategoryEnum Enum with underlying type: string
type ListExcludedObjectsReasonCategoryEnum string

// Set of constants representing the allowable values for ListExcludedObjectsReasonCategoryEnum
const (
	ListExcludedObjectsReasonCategoryOracleMaintained  ListExcludedObjectsReasonCategoryEnum = "ORACLE_MAINTAINED"
	ListExcludedObjectsReasonCategoryGgUnsupported     ListExcludedObjectsReasonCategoryEnum = "GG_UNSUPPORTED"
	ListExcludedObjectsReasonCategoryUserExcluded      ListExcludedObjectsReasonCategoryEnum = "USER_EXCLUDED"
	ListExcludedObjectsReasonCategoryMandatoryExcluded ListExcludedObjectsReasonCategoryEnum = "MANDATORY_EXCLUDED"
	ListExcludedObjectsReasonCategoryUserExcludedType  ListExcludedObjectsReasonCategoryEnum = "USER_EXCLUDED_TYPE"
)

var mappingListExcludedObjectsReasonCategoryEnum = map[string]ListExcludedObjectsReasonCategoryEnum{
	"ORACLE_MAINTAINED":  ListExcludedObjectsReasonCategoryOracleMaintained,
	"GG_UNSUPPORTED":     ListExcludedObjectsReasonCategoryGgUnsupported,
	"USER_EXCLUDED":      ListExcludedObjectsReasonCategoryUserExcluded,
	"MANDATORY_EXCLUDED": ListExcludedObjectsReasonCategoryMandatoryExcluded,
	"USER_EXCLUDED_TYPE": ListExcludedObjectsReasonCategoryUserExcludedType,
}

var mappingListExcludedObjectsReasonCategoryEnumLowerCase = map[string]ListExcludedObjectsReasonCategoryEnum{
	"oracle_maintained":  ListExcludedObjectsReasonCategoryOracleMaintained,
	"gg_unsupported":     ListExcludedObjectsReasonCategoryGgUnsupported,
	"user_excluded":      ListExcludedObjectsReasonCategoryUserExcluded,
	"mandatory_excluded": ListExcludedObjectsReasonCategoryMandatoryExcluded,
	"user_excluded_type": ListExcludedObjectsReasonCategoryUserExcludedType,
}

// GetListExcludedObjectsReasonCategoryEnumValues Enumerates the set of values for ListExcludedObjectsReasonCategoryEnum
func GetListExcludedObjectsReasonCategoryEnumValues() []ListExcludedObjectsReasonCategoryEnum {
	values := make([]ListExcludedObjectsReasonCategoryEnum, 0)
	for _, v := range mappingListExcludedObjectsReasonCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetListExcludedObjectsReasonCategoryEnumStringValues Enumerates the set of values in String for ListExcludedObjectsReasonCategoryEnum
func GetListExcludedObjectsReasonCategoryEnumStringValues() []string {
	return []string{
		"ORACLE_MAINTAINED",
		"GG_UNSUPPORTED",
		"USER_EXCLUDED",
		"MANDATORY_EXCLUDED",
		"USER_EXCLUDED_TYPE",
	}
}

// GetMappingListExcludedObjectsReasonCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListExcludedObjectsReasonCategoryEnum(val string) (ListExcludedObjectsReasonCategoryEnum, bool) {
	enum, ok := mappingListExcludedObjectsReasonCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
