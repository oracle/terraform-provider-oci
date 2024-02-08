// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ValidateAssociationParametersRequest wrapper for the ValidateAssociationParameters operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ValidateAssociationParameters.go.html to see an example of how to use ValidateAssociationParametersRequest.
type ValidateAssociationParametersRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// Details for the new log analytics associations.
	UpsertLogAnalyticsAssociationDetails `contributesTo:"body"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ValidateAssociationParametersSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The attribute used to sort the returned association parameters
	SortBy ValidateAssociationParametersSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ValidateAssociationParametersRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ValidateAssociationParametersRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ValidateAssociationParametersRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ValidateAssociationParametersRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ValidateAssociationParametersRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingValidateAssociationParametersSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetValidateAssociationParametersSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingValidateAssociationParametersSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetValidateAssociationParametersSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ValidateAssociationParametersResponse wrapper for the ValidateAssociationParameters operation
type ValidateAssociationParametersResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogAnalyticsAssociationParameterCollection instances
	LogAnalyticsAssociationParameterCollection `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ValidateAssociationParametersResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ValidateAssociationParametersResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ValidateAssociationParametersSortOrderEnum Enum with underlying type: string
type ValidateAssociationParametersSortOrderEnum string

// Set of constants representing the allowable values for ValidateAssociationParametersSortOrderEnum
const (
	ValidateAssociationParametersSortOrderAsc  ValidateAssociationParametersSortOrderEnum = "ASC"
	ValidateAssociationParametersSortOrderDesc ValidateAssociationParametersSortOrderEnum = "DESC"
)

var mappingValidateAssociationParametersSortOrderEnum = map[string]ValidateAssociationParametersSortOrderEnum{
	"ASC":  ValidateAssociationParametersSortOrderAsc,
	"DESC": ValidateAssociationParametersSortOrderDesc,
}

var mappingValidateAssociationParametersSortOrderEnumLowerCase = map[string]ValidateAssociationParametersSortOrderEnum{
	"asc":  ValidateAssociationParametersSortOrderAsc,
	"desc": ValidateAssociationParametersSortOrderDesc,
}

// GetValidateAssociationParametersSortOrderEnumValues Enumerates the set of values for ValidateAssociationParametersSortOrderEnum
func GetValidateAssociationParametersSortOrderEnumValues() []ValidateAssociationParametersSortOrderEnum {
	values := make([]ValidateAssociationParametersSortOrderEnum, 0)
	for _, v := range mappingValidateAssociationParametersSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetValidateAssociationParametersSortOrderEnumStringValues Enumerates the set of values in String for ValidateAssociationParametersSortOrderEnum
func GetValidateAssociationParametersSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingValidateAssociationParametersSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidateAssociationParametersSortOrderEnum(val string) (ValidateAssociationParametersSortOrderEnum, bool) {
	enum, ok := mappingValidateAssociationParametersSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ValidateAssociationParametersSortByEnum Enum with underlying type: string
type ValidateAssociationParametersSortByEnum string

// Set of constants representing the allowable values for ValidateAssociationParametersSortByEnum
const (
	ValidateAssociationParametersSortBySourcedisplayname ValidateAssociationParametersSortByEnum = "sourceDisplayName"
	ValidateAssociationParametersSortByStatus            ValidateAssociationParametersSortByEnum = "status"
)

var mappingValidateAssociationParametersSortByEnum = map[string]ValidateAssociationParametersSortByEnum{
	"sourceDisplayName": ValidateAssociationParametersSortBySourcedisplayname,
	"status":            ValidateAssociationParametersSortByStatus,
}

var mappingValidateAssociationParametersSortByEnumLowerCase = map[string]ValidateAssociationParametersSortByEnum{
	"sourcedisplayname": ValidateAssociationParametersSortBySourcedisplayname,
	"status":            ValidateAssociationParametersSortByStatus,
}

// GetValidateAssociationParametersSortByEnumValues Enumerates the set of values for ValidateAssociationParametersSortByEnum
func GetValidateAssociationParametersSortByEnumValues() []ValidateAssociationParametersSortByEnum {
	values := make([]ValidateAssociationParametersSortByEnum, 0)
	for _, v := range mappingValidateAssociationParametersSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetValidateAssociationParametersSortByEnumStringValues Enumerates the set of values in String for ValidateAssociationParametersSortByEnum
func GetValidateAssociationParametersSortByEnumStringValues() []string {
	return []string{
		"sourceDisplayName",
		"status",
	}
}

// GetMappingValidateAssociationParametersSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingValidateAssociationParametersSortByEnum(val string) (ValidateAssociationParametersSortByEnum, bool) {
	enum, ok := mappingValidateAssociationParametersSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
