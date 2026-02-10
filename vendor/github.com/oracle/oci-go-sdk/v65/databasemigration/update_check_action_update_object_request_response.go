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

// UpdateCheckActionUpdateObjectRequest wrapper for the UpdateCheckActionUpdateObject operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/databasemigration/UpdateCheckActionUpdateObject.go.html to see an example of how to use UpdateCheckActionUpdateObjectRequest.
type UpdateCheckActionUpdateObjectRequest struct {

	// The OCID of the Assessment
	AssessmentId *string `mandatory:"true" contributesTo:"path" name:"assessmentId"`

	// The name of the Assessor
	AssessorName *string `mandatory:"true" contributesTo:"path" name:"assessorName"`

	// The Name of the assessor check
	CheckName *string `mandatory:"true" contributesTo:"path" name:"checkName"`

	// Collection of AdvisorReportCheckObjectSummary.
	UpdateCheckActionUpdateObjectDetails `contributesTo:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for name is custom based on it's usage frequency. If no value is specified name is default.
	SortBy UpdateCheckActionUpdateObjectSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder UpdateCheckActionUpdateObjectSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request UpdateCheckActionUpdateObjectRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request UpdateCheckActionUpdateObjectRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request UpdateCheckActionUpdateObjectRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request UpdateCheckActionUpdateObjectRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request UpdateCheckActionUpdateObjectRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpdateCheckActionUpdateObjectSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetUpdateCheckActionUpdateObjectSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateCheckActionUpdateObjectSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetUpdateCheckActionUpdateObjectSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateCheckActionUpdateObjectResponse wrapper for the UpdateCheckActionUpdateObject operation
type UpdateCheckActionUpdateObjectResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response UpdateCheckActionUpdateObjectResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response UpdateCheckActionUpdateObjectResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// UpdateCheckActionUpdateObjectSortByEnum Enum with underlying type: string
type UpdateCheckActionUpdateObjectSortByEnum string

// Set of constants representing the allowable values for UpdateCheckActionUpdateObjectSortByEnum
const (
	UpdateCheckActionUpdateObjectSortByName UpdateCheckActionUpdateObjectSortByEnum = "name"
)

var mappingUpdateCheckActionUpdateObjectSortByEnum = map[string]UpdateCheckActionUpdateObjectSortByEnum{
	"name": UpdateCheckActionUpdateObjectSortByName,
}

var mappingUpdateCheckActionUpdateObjectSortByEnumLowerCase = map[string]UpdateCheckActionUpdateObjectSortByEnum{
	"name": UpdateCheckActionUpdateObjectSortByName,
}

// GetUpdateCheckActionUpdateObjectSortByEnumValues Enumerates the set of values for UpdateCheckActionUpdateObjectSortByEnum
func GetUpdateCheckActionUpdateObjectSortByEnumValues() []UpdateCheckActionUpdateObjectSortByEnum {
	values := make([]UpdateCheckActionUpdateObjectSortByEnum, 0)
	for _, v := range mappingUpdateCheckActionUpdateObjectSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCheckActionUpdateObjectSortByEnumStringValues Enumerates the set of values in String for UpdateCheckActionUpdateObjectSortByEnum
func GetUpdateCheckActionUpdateObjectSortByEnumStringValues() []string {
	return []string{
		"name",
	}
}

// GetMappingUpdateCheckActionUpdateObjectSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCheckActionUpdateObjectSortByEnum(val string) (UpdateCheckActionUpdateObjectSortByEnum, bool) {
	enum, ok := mappingUpdateCheckActionUpdateObjectSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdateCheckActionUpdateObjectSortOrderEnum Enum with underlying type: string
type UpdateCheckActionUpdateObjectSortOrderEnum string

// Set of constants representing the allowable values for UpdateCheckActionUpdateObjectSortOrderEnum
const (
	UpdateCheckActionUpdateObjectSortOrderAsc  UpdateCheckActionUpdateObjectSortOrderEnum = "ASC"
	UpdateCheckActionUpdateObjectSortOrderDesc UpdateCheckActionUpdateObjectSortOrderEnum = "DESC"
)

var mappingUpdateCheckActionUpdateObjectSortOrderEnum = map[string]UpdateCheckActionUpdateObjectSortOrderEnum{
	"ASC":  UpdateCheckActionUpdateObjectSortOrderAsc,
	"DESC": UpdateCheckActionUpdateObjectSortOrderDesc,
}

var mappingUpdateCheckActionUpdateObjectSortOrderEnumLowerCase = map[string]UpdateCheckActionUpdateObjectSortOrderEnum{
	"asc":  UpdateCheckActionUpdateObjectSortOrderAsc,
	"desc": UpdateCheckActionUpdateObjectSortOrderDesc,
}

// GetUpdateCheckActionUpdateObjectSortOrderEnumValues Enumerates the set of values for UpdateCheckActionUpdateObjectSortOrderEnum
func GetUpdateCheckActionUpdateObjectSortOrderEnumValues() []UpdateCheckActionUpdateObjectSortOrderEnum {
	values := make([]UpdateCheckActionUpdateObjectSortOrderEnum, 0)
	for _, v := range mappingUpdateCheckActionUpdateObjectSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateCheckActionUpdateObjectSortOrderEnumStringValues Enumerates the set of values in String for UpdateCheckActionUpdateObjectSortOrderEnum
func GetUpdateCheckActionUpdateObjectSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingUpdateCheckActionUpdateObjectSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateCheckActionUpdateObjectSortOrderEnum(val string) (UpdateCheckActionUpdateObjectSortOrderEnum, bool) {
	enum, ok := mappingUpdateCheckActionUpdateObjectSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
