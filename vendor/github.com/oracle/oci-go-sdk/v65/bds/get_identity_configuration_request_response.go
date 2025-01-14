// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// GetIdentityConfigurationRequest wrapper for the GetIdentityConfiguration operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/bds/GetIdentityConfiguration.go.html to see an example of how to use GetIdentityConfigurationRequest.
type GetIdentityConfigurationRequest struct {

	// The OCID of the cluster.
	BdsInstanceId *string `mandatory:"true" contributesTo:"path" name:"bdsInstanceId"`

	// The OCID of the identity configuration
	IdentityConfigurationId *string `mandatory:"true" contributesTo:"path" name:"identityConfigurationId"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy GetIdentityConfigurationSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder GetIdentityConfigurationSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request GetIdentityConfigurationRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetIdentityConfigurationRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetIdentityConfigurationRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetIdentityConfigurationRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetIdentityConfigurationRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetIdentityConfigurationSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetGetIdentityConfigurationSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingGetIdentityConfigurationSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetGetIdentityConfigurationSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetIdentityConfigurationResponse wrapper for the GetIdentityConfiguration operation
type GetIdentityConfigurationResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of IdentityConfiguration instances
	IdentityConfiguration `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a request, provide this request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`
}

func (response GetIdentityConfigurationResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetIdentityConfigurationResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetIdentityConfigurationSortByEnum Enum with underlying type: string
type GetIdentityConfigurationSortByEnum string

// Set of constants representing the allowable values for GetIdentityConfigurationSortByEnum
const (
	GetIdentityConfigurationSortByTimecreated GetIdentityConfigurationSortByEnum = "timeCreated"
	GetIdentityConfigurationSortByDisplayname GetIdentityConfigurationSortByEnum = "displayName"
)

var mappingGetIdentityConfigurationSortByEnum = map[string]GetIdentityConfigurationSortByEnum{
	"timeCreated": GetIdentityConfigurationSortByTimecreated,
	"displayName": GetIdentityConfigurationSortByDisplayname,
}

var mappingGetIdentityConfigurationSortByEnumLowerCase = map[string]GetIdentityConfigurationSortByEnum{
	"timecreated": GetIdentityConfigurationSortByTimecreated,
	"displayname": GetIdentityConfigurationSortByDisplayname,
}

// GetGetIdentityConfigurationSortByEnumValues Enumerates the set of values for GetIdentityConfigurationSortByEnum
func GetGetIdentityConfigurationSortByEnumValues() []GetIdentityConfigurationSortByEnum {
	values := make([]GetIdentityConfigurationSortByEnum, 0)
	for _, v := range mappingGetIdentityConfigurationSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetGetIdentityConfigurationSortByEnumStringValues Enumerates the set of values in String for GetIdentityConfigurationSortByEnum
func GetGetIdentityConfigurationSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingGetIdentityConfigurationSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetIdentityConfigurationSortByEnum(val string) (GetIdentityConfigurationSortByEnum, bool) {
	enum, ok := mappingGetIdentityConfigurationSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// GetIdentityConfigurationSortOrderEnum Enum with underlying type: string
type GetIdentityConfigurationSortOrderEnum string

// Set of constants representing the allowable values for GetIdentityConfigurationSortOrderEnum
const (
	GetIdentityConfigurationSortOrderAsc  GetIdentityConfigurationSortOrderEnum = "ASC"
	GetIdentityConfigurationSortOrderDesc GetIdentityConfigurationSortOrderEnum = "DESC"
)

var mappingGetIdentityConfigurationSortOrderEnum = map[string]GetIdentityConfigurationSortOrderEnum{
	"ASC":  GetIdentityConfigurationSortOrderAsc,
	"DESC": GetIdentityConfigurationSortOrderDesc,
}

var mappingGetIdentityConfigurationSortOrderEnumLowerCase = map[string]GetIdentityConfigurationSortOrderEnum{
	"asc":  GetIdentityConfigurationSortOrderAsc,
	"desc": GetIdentityConfigurationSortOrderDesc,
}

// GetGetIdentityConfigurationSortOrderEnumValues Enumerates the set of values for GetIdentityConfigurationSortOrderEnum
func GetGetIdentityConfigurationSortOrderEnumValues() []GetIdentityConfigurationSortOrderEnum {
	values := make([]GetIdentityConfigurationSortOrderEnum, 0)
	for _, v := range mappingGetIdentityConfigurationSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetGetIdentityConfigurationSortOrderEnumStringValues Enumerates the set of values in String for GetIdentityConfigurationSortOrderEnum
func GetGetIdentityConfigurationSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingGetIdentityConfigurationSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetIdentityConfigurationSortOrderEnum(val string) (GetIdentityConfigurationSortOrderEnum, bool) {
	enum, ok := mappingGetIdentityConfigurationSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
