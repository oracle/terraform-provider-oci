// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datascience

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListModelArtifactSignaturesRequest wrapper for the ListModelArtifactSignatures operation
type ListModelArtifactSignaturesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
	ModelId *string `mandatory:"true" contributesTo:"path" name:"modelId"`

	// <b>Filter</b> results by the OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// <b>Filter</b> results by its user-friendly name.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The digest of the model artifact.
	Digest *string `mandatory:"false" contributesTo:"query" name:"digest"`

	// <b>Filter</b> results by the specified lifecycle state. Must be a valid
	// state for the resource type.
	LifecycleState ListModelArtifactSignaturesLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 100 is the maximum.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// See List Pagination (https://docs.oracle.com/iaas/Content/General/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Specifies sort order to use, either `ASC` (ascending) or `DESC` (descending).
	SortOrder ListModelArtifactSignaturesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Specifies the field to sort by. Accepts only one field.
	// By default, when you sort by `timeCreated`, the results are shown
	// in descending order. All other fields default to ascending order. Sort order for the `displayName` field is case sensitive.
	SortBy ListModelArtifactSignaturesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle assigned identifier for the request. If you need to contact Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListModelArtifactSignaturesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListModelArtifactSignaturesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListModelArtifactSignaturesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListModelArtifactSignaturesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListModelArtifactSignaturesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListModelArtifactSignaturesLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListModelArtifactSignaturesLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelArtifactSignaturesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListModelArtifactSignaturesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListModelArtifactSignaturesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListModelArtifactSignaturesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListModelArtifactSignaturesResponse wrapper for the ListModelArtifactSignatures operation
type ListModelArtifactSignaturesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ModelArtifactSignatureSummary instances
	Items []ModelArtifactSignatureSummary `presentIn:"body"`

	// Retrieves the next page of results. When this header appears in the response, additional pages of results remain. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Retrieves the previous page of results. When this header appears in the response, previous pages of results exist. See List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`

	// Unique Oracle assigned identifier for the request. If you need to contact
	// Oracle about a particular request, then provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListModelArtifactSignaturesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListModelArtifactSignaturesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListModelArtifactSignaturesLifecycleStateEnum Enum with underlying type: string
type ListModelArtifactSignaturesLifecycleStateEnum string

// Set of constants representing the allowable values for ListModelArtifactSignaturesLifecycleStateEnum
const (
	ListModelArtifactSignaturesLifecycleStateActive   ListModelArtifactSignaturesLifecycleStateEnum = "ACTIVE"
	ListModelArtifactSignaturesLifecycleStateCreating ListModelArtifactSignaturesLifecycleStateEnum = "CREATING"
	ListModelArtifactSignaturesLifecycleStateDeleted  ListModelArtifactSignaturesLifecycleStateEnum = "DELETED"
	ListModelArtifactSignaturesLifecycleStateFailed   ListModelArtifactSignaturesLifecycleStateEnum = "FAILED"
)

var mappingListModelArtifactSignaturesLifecycleStateEnum = map[string]ListModelArtifactSignaturesLifecycleStateEnum{
	"ACTIVE":   ListModelArtifactSignaturesLifecycleStateActive,
	"CREATING": ListModelArtifactSignaturesLifecycleStateCreating,
	"DELETED":  ListModelArtifactSignaturesLifecycleStateDeleted,
	"FAILED":   ListModelArtifactSignaturesLifecycleStateFailed,
}

var mappingListModelArtifactSignaturesLifecycleStateEnumLowerCase = map[string]ListModelArtifactSignaturesLifecycleStateEnum{
	"active":   ListModelArtifactSignaturesLifecycleStateActive,
	"creating": ListModelArtifactSignaturesLifecycleStateCreating,
	"deleted":  ListModelArtifactSignaturesLifecycleStateDeleted,
	"failed":   ListModelArtifactSignaturesLifecycleStateFailed,
}

// GetListModelArtifactSignaturesLifecycleStateEnumValues Enumerates the set of values for ListModelArtifactSignaturesLifecycleStateEnum
func GetListModelArtifactSignaturesLifecycleStateEnumValues() []ListModelArtifactSignaturesLifecycleStateEnum {
	values := make([]ListModelArtifactSignaturesLifecycleStateEnum, 0)
	for _, v := range mappingListModelArtifactSignaturesLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelArtifactSignaturesLifecycleStateEnumStringValues Enumerates the set of values in String for ListModelArtifactSignaturesLifecycleStateEnum
func GetListModelArtifactSignaturesLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingListModelArtifactSignaturesLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelArtifactSignaturesLifecycleStateEnum(val string) (ListModelArtifactSignaturesLifecycleStateEnum, bool) {
	enum, ok := mappingListModelArtifactSignaturesLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelArtifactSignaturesSortOrderEnum Enum with underlying type: string
type ListModelArtifactSignaturesSortOrderEnum string

// Set of constants representing the allowable values for ListModelArtifactSignaturesSortOrderEnum
const (
	ListModelArtifactSignaturesSortOrderAsc  ListModelArtifactSignaturesSortOrderEnum = "ASC"
	ListModelArtifactSignaturesSortOrderDesc ListModelArtifactSignaturesSortOrderEnum = "DESC"
)

var mappingListModelArtifactSignaturesSortOrderEnum = map[string]ListModelArtifactSignaturesSortOrderEnum{
	"ASC":  ListModelArtifactSignaturesSortOrderAsc,
	"DESC": ListModelArtifactSignaturesSortOrderDesc,
}

var mappingListModelArtifactSignaturesSortOrderEnumLowerCase = map[string]ListModelArtifactSignaturesSortOrderEnum{
	"asc":  ListModelArtifactSignaturesSortOrderAsc,
	"desc": ListModelArtifactSignaturesSortOrderDesc,
}

// GetListModelArtifactSignaturesSortOrderEnumValues Enumerates the set of values for ListModelArtifactSignaturesSortOrderEnum
func GetListModelArtifactSignaturesSortOrderEnumValues() []ListModelArtifactSignaturesSortOrderEnum {
	values := make([]ListModelArtifactSignaturesSortOrderEnum, 0)
	for _, v := range mappingListModelArtifactSignaturesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelArtifactSignaturesSortOrderEnumStringValues Enumerates the set of values in String for ListModelArtifactSignaturesSortOrderEnum
func GetListModelArtifactSignaturesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListModelArtifactSignaturesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelArtifactSignaturesSortOrderEnum(val string) (ListModelArtifactSignaturesSortOrderEnum, bool) {
	enum, ok := mappingListModelArtifactSignaturesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListModelArtifactSignaturesSortByEnum Enum with underlying type: string
type ListModelArtifactSignaturesSortByEnum string

// Set of constants representing the allowable values for ListModelArtifactSignaturesSortByEnum
const (
	ListModelArtifactSignaturesSortByTimecreated    ListModelArtifactSignaturesSortByEnum = "timeCreated"
	ListModelArtifactSignaturesSortByDisplayname    ListModelArtifactSignaturesSortByEnum = "displayName"
	ListModelArtifactSignaturesSortByLifecyclestate ListModelArtifactSignaturesSortByEnum = "lifecycleState"
)

var mappingListModelArtifactSignaturesSortByEnum = map[string]ListModelArtifactSignaturesSortByEnum{
	"timeCreated":    ListModelArtifactSignaturesSortByTimecreated,
	"displayName":    ListModelArtifactSignaturesSortByDisplayname,
	"lifecycleState": ListModelArtifactSignaturesSortByLifecyclestate,
}

var mappingListModelArtifactSignaturesSortByEnumLowerCase = map[string]ListModelArtifactSignaturesSortByEnum{
	"timecreated":    ListModelArtifactSignaturesSortByTimecreated,
	"displayname":    ListModelArtifactSignaturesSortByDisplayname,
	"lifecyclestate": ListModelArtifactSignaturesSortByLifecyclestate,
}

// GetListModelArtifactSignaturesSortByEnumValues Enumerates the set of values for ListModelArtifactSignaturesSortByEnum
func GetListModelArtifactSignaturesSortByEnumValues() []ListModelArtifactSignaturesSortByEnum {
	values := make([]ListModelArtifactSignaturesSortByEnum, 0)
	for _, v := range mappingListModelArtifactSignaturesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListModelArtifactSignaturesSortByEnumStringValues Enumerates the set of values in String for ListModelArtifactSignaturesSortByEnum
func GetListModelArtifactSignaturesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
		"lifecycleState",
	}
}

// GetMappingListModelArtifactSignaturesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListModelArtifactSignaturesSortByEnum(val string) (ListModelArtifactSignaturesSortByEnum, bool) {
	enum, ok := mappingListModelArtifactSignaturesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
