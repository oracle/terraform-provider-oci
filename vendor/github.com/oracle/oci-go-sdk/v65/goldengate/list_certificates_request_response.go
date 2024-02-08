// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCertificatesRequest wrapper for the ListCertificates operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/goldengate/ListCertificates.go.html to see an example of how to use ListCertificatesRequest.
type ListCertificatesRequest struct {

	// A unique Deployment identifier.
	DeploymentId *string `mandatory:"true" contributesTo:"path" name:"deploymentId"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only connections having the 'lifecycleState' given.
	LifecycleState ListCertificatesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually
	// retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'asc' or 'desc'.
	SortOrder ListCertificatesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order can be provided. Default order for 'timeCreated' is
	// descending.  Default order for 'displayName' is ascending. If no value is specified
	// timeCreated is the default.
	SortBy ListCertificatesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCertificatesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCertificatesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCertificatesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCertificatesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCertificatesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListCertificatesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListCertificatesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificatesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCertificatesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCertificatesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCertificatesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCertificatesResponse wrapper for the ListCertificates operation
type ListCertificatesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CertificateCollection instances
	CertificateCollection `presentIn:"body"`

	// A unique Oracle-assigned identifier for the request. If you need to contact Oracle about a
	// particular request, please include the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// The page token represents the page to start retrieving results. This is usually retrieved
	// from a previous list call.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCertificatesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCertificatesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCertificatesLifecycleStateEnum Enum with underlying type: string
type ListCertificatesLifecycleStateEnum string

// Set of constants representing the allowable values for ListCertificatesLifecycleStateEnum
const (
	ListCertificatesLifecycleStateCreating ListCertificatesLifecycleStateEnum = "CREATING"
	ListCertificatesLifecycleStateActive   ListCertificatesLifecycleStateEnum = "ACTIVE"
	ListCertificatesLifecycleStateDeleting ListCertificatesLifecycleStateEnum = "DELETING"
	ListCertificatesLifecycleStateDeleted  ListCertificatesLifecycleStateEnum = "DELETED"
	ListCertificatesLifecycleStateFailed   ListCertificatesLifecycleStateEnum = "FAILED"
)

var mappingListCertificatesLifecycleStateEnum = map[string]ListCertificatesLifecycleStateEnum{
	"CREATING": ListCertificatesLifecycleStateCreating,
	"ACTIVE":   ListCertificatesLifecycleStateActive,
	"DELETING": ListCertificatesLifecycleStateDeleting,
	"DELETED":  ListCertificatesLifecycleStateDeleted,
	"FAILED":   ListCertificatesLifecycleStateFailed,
}

var mappingListCertificatesLifecycleStateEnumLowerCase = map[string]ListCertificatesLifecycleStateEnum{
	"creating": ListCertificatesLifecycleStateCreating,
	"active":   ListCertificatesLifecycleStateActive,
	"deleting": ListCertificatesLifecycleStateDeleting,
	"deleted":  ListCertificatesLifecycleStateDeleted,
	"failed":   ListCertificatesLifecycleStateFailed,
}

// GetListCertificatesLifecycleStateEnumValues Enumerates the set of values for ListCertificatesLifecycleStateEnum
func GetListCertificatesLifecycleStateEnumValues() []ListCertificatesLifecycleStateEnum {
	values := make([]ListCertificatesLifecycleStateEnum, 0)
	for _, v := range mappingListCertificatesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificatesLifecycleStateEnumStringValues Enumerates the set of values in String for ListCertificatesLifecycleStateEnum
func GetListCertificatesLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListCertificatesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificatesLifecycleStateEnum(val string) (ListCertificatesLifecycleStateEnum, bool) {
	enum, ok := mappingListCertificatesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificatesSortOrderEnum Enum with underlying type: string
type ListCertificatesSortOrderEnum string

// Set of constants representing the allowable values for ListCertificatesSortOrderEnum
const (
	ListCertificatesSortOrderAsc  ListCertificatesSortOrderEnum = "ASC"
	ListCertificatesSortOrderDesc ListCertificatesSortOrderEnum = "DESC"
)

var mappingListCertificatesSortOrderEnum = map[string]ListCertificatesSortOrderEnum{
	"ASC":  ListCertificatesSortOrderAsc,
	"DESC": ListCertificatesSortOrderDesc,
}

var mappingListCertificatesSortOrderEnumLowerCase = map[string]ListCertificatesSortOrderEnum{
	"asc":  ListCertificatesSortOrderAsc,
	"desc": ListCertificatesSortOrderDesc,
}

// GetListCertificatesSortOrderEnumValues Enumerates the set of values for ListCertificatesSortOrderEnum
func GetListCertificatesSortOrderEnumValues() []ListCertificatesSortOrderEnum {
	values := make([]ListCertificatesSortOrderEnum, 0)
	for _, v := range mappingListCertificatesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificatesSortOrderEnumStringValues Enumerates the set of values in String for ListCertificatesSortOrderEnum
func GetListCertificatesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCertificatesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificatesSortOrderEnum(val string) (ListCertificatesSortOrderEnum, bool) {
	enum, ok := mappingListCertificatesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCertificatesSortByEnum Enum with underlying type: string
type ListCertificatesSortByEnum string

// Set of constants representing the allowable values for ListCertificatesSortByEnum
const (
	ListCertificatesSortByTimecreated ListCertificatesSortByEnum = "timeCreated"
	ListCertificatesSortByDisplayname ListCertificatesSortByEnum = "displayName"
)

var mappingListCertificatesSortByEnum = map[string]ListCertificatesSortByEnum{
	"timeCreated": ListCertificatesSortByTimecreated,
	"displayName": ListCertificatesSortByDisplayname,
}

var mappingListCertificatesSortByEnumLowerCase = map[string]ListCertificatesSortByEnum{
	"timecreated": ListCertificatesSortByTimecreated,
	"displayname": ListCertificatesSortByDisplayname,
}

// GetListCertificatesSortByEnumValues Enumerates the set of values for ListCertificatesSortByEnum
func GetListCertificatesSortByEnumValues() []ListCertificatesSortByEnum {
	values := make([]ListCertificatesSortByEnum, 0)
	for _, v := range mappingListCertificatesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCertificatesSortByEnumStringValues Enumerates the set of values in String for ListCertificatesSortByEnum
func GetListCertificatesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCertificatesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCertificatesSortByEnum(val string) (ListCertificatesSortByEnum, bool) {
	enum, ok := mappingListCertificatesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
