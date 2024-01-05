// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListAwrHubSourcesRequest wrapper for the ListAwrHubSources operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/opsi/ListAwrHubSources.go.html to see an example of how to use ListAwrHubSourcesRequest.
type ListAwrHubSourcesRequest struct {

	// Unique Awr Hub identifier
	AwrHubId *string `mandatory:"true" contributesTo:"query" name:"awrHubId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// Awr Hub source identifier
	AwrHubSourceId *string `mandatory:"false" contributesTo:"query" name:"awrHubSourceId"`

	// Filter by one or more database type.
	// Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
	SourceType []AwrHubSourceTypeEnum `contributesTo:"query" name:"sourceType" omitEmpty:"true" collectionFormat:"multi"`

	// Awr Hub source database name
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// Resource Status
	Status []AwrHubSourceStatusEnum `contributesTo:"query" name:"status" omitEmpty:"true" collectionFormat:"multi"`

	// Lifecycle states
	LifecycleState []AwrHubSourceLifecycleStateEnum `contributesTo:"query" name:"lifecycleState" omitEmpty:"true" collectionFormat:"multi"`

	// For list pagination. The maximum number of results per page, or items to
	// return in a paginated "List" call.
	// For important details about how pagination works, see
	// List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	// Example: `50`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response header from
	// the previous "List" call. For important details about how pagination works,
	// see List Pagination (https://docs.cloud.oracle.com/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListAwrHubSourcesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending. If no value is specified timeCreated is default.
	SortBy ListAwrHubSourcesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListAwrHubSourcesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListAwrHubSourcesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListAwrHubSourcesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListAwrHubSourcesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListAwrHubSourcesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	for _, val := range request.SourceType {
		if _, ok := GetMappingAwrHubSourceTypeEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", val, strings.Join(GetAwrHubSourceTypeEnumStringValues(), ",")))
		}
	}

	for _, val := range request.Status {
		if _, ok := GetMappingAwrHubSourceStatusEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", val, strings.Join(GetAwrHubSourceStatusEnumStringValues(), ",")))
		}
	}

	for _, val := range request.LifecycleState {
		if _, ok := GetMappingAwrHubSourceLifecycleStateEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", val, strings.Join(GetAwrHubSourceLifecycleStateEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingListAwrHubSourcesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListAwrHubSourcesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListAwrHubSourcesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListAwrHubSourcesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListAwrHubSourcesResponse wrapper for the ListAwrHubSources operation
type ListAwrHubSourcesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of AwrHubSourceSummaryCollection instances
	AwrHubSourceSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListAwrHubSourcesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListAwrHubSourcesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListAwrHubSourcesSortOrderEnum Enum with underlying type: string
type ListAwrHubSourcesSortOrderEnum string

// Set of constants representing the allowable values for ListAwrHubSourcesSortOrderEnum
const (
	ListAwrHubSourcesSortOrderAsc  ListAwrHubSourcesSortOrderEnum = "ASC"
	ListAwrHubSourcesSortOrderDesc ListAwrHubSourcesSortOrderEnum = "DESC"
)

var mappingListAwrHubSourcesSortOrderEnum = map[string]ListAwrHubSourcesSortOrderEnum{
	"ASC":  ListAwrHubSourcesSortOrderAsc,
	"DESC": ListAwrHubSourcesSortOrderDesc,
}

var mappingListAwrHubSourcesSortOrderEnumLowerCase = map[string]ListAwrHubSourcesSortOrderEnum{
	"asc":  ListAwrHubSourcesSortOrderAsc,
	"desc": ListAwrHubSourcesSortOrderDesc,
}

// GetListAwrHubSourcesSortOrderEnumValues Enumerates the set of values for ListAwrHubSourcesSortOrderEnum
func GetListAwrHubSourcesSortOrderEnumValues() []ListAwrHubSourcesSortOrderEnum {
	values := make([]ListAwrHubSourcesSortOrderEnum, 0)
	for _, v := range mappingListAwrHubSourcesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrHubSourcesSortOrderEnumStringValues Enumerates the set of values in String for ListAwrHubSourcesSortOrderEnum
func GetListAwrHubSourcesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListAwrHubSourcesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrHubSourcesSortOrderEnum(val string) (ListAwrHubSourcesSortOrderEnum, bool) {
	enum, ok := mappingListAwrHubSourcesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListAwrHubSourcesSortByEnum Enum with underlying type: string
type ListAwrHubSourcesSortByEnum string

// Set of constants representing the allowable values for ListAwrHubSourcesSortByEnum
const (
	ListAwrHubSourcesSortByTimecreated ListAwrHubSourcesSortByEnum = "timeCreated"
	ListAwrHubSourcesSortByDisplayname ListAwrHubSourcesSortByEnum = "displayName"
)

var mappingListAwrHubSourcesSortByEnum = map[string]ListAwrHubSourcesSortByEnum{
	"timeCreated": ListAwrHubSourcesSortByTimecreated,
	"displayName": ListAwrHubSourcesSortByDisplayname,
}

var mappingListAwrHubSourcesSortByEnumLowerCase = map[string]ListAwrHubSourcesSortByEnum{
	"timecreated": ListAwrHubSourcesSortByTimecreated,
	"displayname": ListAwrHubSourcesSortByDisplayname,
}

// GetListAwrHubSourcesSortByEnumValues Enumerates the set of values for ListAwrHubSourcesSortByEnum
func GetListAwrHubSourcesSortByEnumValues() []ListAwrHubSourcesSortByEnum {
	values := make([]ListAwrHubSourcesSortByEnum, 0)
	for _, v := range mappingListAwrHubSourcesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListAwrHubSourcesSortByEnumStringValues Enumerates the set of values in String for ListAwrHubSourcesSortByEnum
func GetListAwrHubSourcesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListAwrHubSourcesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListAwrHubSourcesSortByEnum(val string) (ListAwrHubSourcesSortByEnum, bool) {
	enum, ok := mappingListAwrHubSourcesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
