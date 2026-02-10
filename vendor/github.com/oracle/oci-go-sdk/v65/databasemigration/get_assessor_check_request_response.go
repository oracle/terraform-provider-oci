// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetAssessorCheckRequest wrapper for the GetAssessorCheck operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/GetAssessorCheck.go.html to see an example of how to use GetAssessorCheckRequest.
type GetAssessorCheckRequest struct {

	// The OCID of the Assessment
	AssessmentId *string `mandatory:"true" contributesTo:"path" name:"assessmentId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the Assessor
	AssessorName *string `mandatory:"true" contributesTo:"path" name:"assessorName"`

	// The Name of the assessor check
	CheckName *string `mandatory:"true" contributesTo:"path" name:"checkName"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending.
	// Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy GetAssessorCheckSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetAssessorCheckSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetAssessorCheckRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetAssessorCheckRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetAssessorCheckRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetAssessorCheckRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetAssessorCheckRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetAssessorCheckSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetAssessorCheckSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetAssessorCheckSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetAssessorCheckSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetAssessorCheckResponse wrapper for the GetAssessorCheck operation
type GetAssessorCheckResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AssessorCheck instances
	AssessorCheck `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response GetAssessorCheckResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetAssessorCheckResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetAssessorCheckSortByEnum Enum with underlying type: string
type GetAssessorCheckSortByEnum string

// Set of constants representing the allowable values for GetAssessorCheckSortByEnum
const (
	GetAssessorCheckSortByTimecreated GetAssessorCheckSortByEnum = "timeCreated"
	GetAssessorCheckSortByDisplayname GetAssessorCheckSortByEnum = "displayName"
)

var mappingGetAssessorCheckSortByEnum = map[string]GetAssessorCheckSortByEnum{
	"timeCreated": GetAssessorCheckSortByTimecreated,
	"displayName": GetAssessorCheckSortByDisplayname,
}

var mappingGetAssessorCheckSortByEnumLowerCase = map[string]GetAssessorCheckSortByEnum{
	"timecreated": GetAssessorCheckSortByTimecreated,
	"displayname": GetAssessorCheckSortByDisplayname,
}

// GetGetAssessorCheckSortByEnumValues Enumerates the set of values for GetAssessorCheckSortByEnum
func GetGetAssessorCheckSortByEnumValues() []GetAssessorCheckSortByEnum {
	values := make([]GetAssessorCheckSortByEnum, 0)
	for _, v := range mappingGetAssessorCheckSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAssessorCheckSortByEnumStringValues Enumerates the set of values in String for GetAssessorCheckSortByEnum
func GetGetAssessorCheckSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingGetAssessorCheckSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAssessorCheckSortByEnum(val string) (GetAssessorCheckSortByEnum, bool) {
	enum, ok := mappingGetAssessorCheckSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetAssessorCheckSortOrderEnum Enum with underlying type: string
type GetAssessorCheckSortOrderEnum string

// Set of constants representing the allowable values for GetAssessorCheckSortOrderEnum
const (
	GetAssessorCheckSortOrderAsc  GetAssessorCheckSortOrderEnum = "ASC"
	GetAssessorCheckSortOrderDesc GetAssessorCheckSortOrderEnum = "DESC"
)

var mappingGetAssessorCheckSortOrderEnum = map[string]GetAssessorCheckSortOrderEnum{
	"ASC":  GetAssessorCheckSortOrderAsc,
	"DESC": GetAssessorCheckSortOrderDesc,
}

var mappingGetAssessorCheckSortOrderEnumLowerCase = map[string]GetAssessorCheckSortOrderEnum{
	"asc":  GetAssessorCheckSortOrderAsc,
	"desc": GetAssessorCheckSortOrderDesc,
}

// GetGetAssessorCheckSortOrderEnumValues Enumerates the set of values for GetAssessorCheckSortOrderEnum
func GetGetAssessorCheckSortOrderEnumValues() []GetAssessorCheckSortOrderEnum {
	values := make([]GetAssessorCheckSortOrderEnum, 0)
	for _, v := range mappingGetAssessorCheckSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetAssessorCheckSortOrderEnumStringValues Enumerates the set of values in String for GetAssessorCheckSortOrderEnum
func GetGetAssessorCheckSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetAssessorCheckSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetAssessorCheckSortOrderEnum(val string) (GetAssessorCheckSortOrderEnum, bool) {
	enum, ok := mappingGetAssessorCheckSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
