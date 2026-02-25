// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFsuReadinessChecksRequest wrapper for the ListFsuReadinessChecks operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuReadinessChecks.go.html to see an example of how to use ListFsuReadinessChecksRequest.
type ListFsuReadinessChecksRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource related to the Exadata Fleet Update Readiness Check.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// A filter to return only resources whose lifecycleState matches the specified lifecycleState.
	LifecycleState FsuReadinessCheckLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources whose type matches the specified type.
	Type FsuReadinessCheckTypeEnum `mandatory:"false" contributesTo:"query" name:"type" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuReadinessChecksSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListFsuReadinessChecksSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuReadinessChecksRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuReadinessChecksRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuReadinessChecksRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuReadinessChecksRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuReadinessChecksRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFsuReadinessCheckLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetFsuReadinessCheckLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFsuReadinessCheckTypeEnum(string(request.Type)); !ok && request.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", request.Type, strings.Join(GetFsuReadinessCheckTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuReadinessChecksSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuReadinessChecksSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuReadinessChecksSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuReadinessChecksSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuReadinessChecksResponse wrapper for the ListFsuReadinessChecks operation
type ListFsuReadinessChecksResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of FsuReadinessCheckCollection instances
	FsuReadinessCheckCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuReadinessChecksResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuReadinessChecksResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuReadinessChecksSortOrderEnum Enum with underlying type: string
type ListFsuReadinessChecksSortOrderEnum string

// Set of constants representing the allowable values for ListFsuReadinessChecksSortOrderEnum
const (
	ListFsuReadinessChecksSortOrderAsc  ListFsuReadinessChecksSortOrderEnum = "ASC"
	ListFsuReadinessChecksSortOrderDesc ListFsuReadinessChecksSortOrderEnum = "DESC"
)

var mappingListFsuReadinessChecksSortOrderEnum = map[string]ListFsuReadinessChecksSortOrderEnum{
	"ASC":  ListFsuReadinessChecksSortOrderAsc,
	"DESC": ListFsuReadinessChecksSortOrderDesc,
}

var mappingListFsuReadinessChecksSortOrderEnumLowerCase = map[string]ListFsuReadinessChecksSortOrderEnum{
	"asc":  ListFsuReadinessChecksSortOrderAsc,
	"desc": ListFsuReadinessChecksSortOrderDesc,
}

// GetListFsuReadinessChecksSortOrderEnumValues Enumerates the set of values for ListFsuReadinessChecksSortOrderEnum
func GetListFsuReadinessChecksSortOrderEnumValues() []ListFsuReadinessChecksSortOrderEnum {
	values := make([]ListFsuReadinessChecksSortOrderEnum, 0)
	for _, v := range mappingListFsuReadinessChecksSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuReadinessChecksSortOrderEnumStringValues Enumerates the set of values in String for ListFsuReadinessChecksSortOrderEnum
func GetListFsuReadinessChecksSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuReadinessChecksSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuReadinessChecksSortOrderEnum(val string) (ListFsuReadinessChecksSortOrderEnum, bool) {
	enum, ok := mappingListFsuReadinessChecksSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuReadinessChecksSortByEnum Enum with underlying type: string
type ListFsuReadinessChecksSortByEnum string

// Set of constants representing the allowable values for ListFsuReadinessChecksSortByEnum
const (
	ListFsuReadinessChecksSortByTimecreated ListFsuReadinessChecksSortByEnum = "timeCreated"
	ListFsuReadinessChecksSortByDisplayname ListFsuReadinessChecksSortByEnum = "displayName"
)

var mappingListFsuReadinessChecksSortByEnum = map[string]ListFsuReadinessChecksSortByEnum{
	"timeCreated": ListFsuReadinessChecksSortByTimecreated,
	"displayName": ListFsuReadinessChecksSortByDisplayname,
}

var mappingListFsuReadinessChecksSortByEnumLowerCase = map[string]ListFsuReadinessChecksSortByEnum{
	"timecreated": ListFsuReadinessChecksSortByTimecreated,
	"displayname": ListFsuReadinessChecksSortByDisplayname,
}

// GetListFsuReadinessChecksSortByEnumValues Enumerates the set of values for ListFsuReadinessChecksSortByEnum
func GetListFsuReadinessChecksSortByEnumValues() []ListFsuReadinessChecksSortByEnum {
	values := make([]ListFsuReadinessChecksSortByEnum, 0)
	for _, v := range mappingListFsuReadinessChecksSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuReadinessChecksSortByEnumStringValues Enumerates the set of values in String for ListFsuReadinessChecksSortByEnum
func GetListFsuReadinessChecksSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListFsuReadinessChecksSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuReadinessChecksSortByEnum(val string) (ListFsuReadinessChecksSortByEnum, bool) {
	enum, ok := mappingListFsuReadinessChecksSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
