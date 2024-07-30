// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetsoftwareupdate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFsuDiscoveryTargetsRequest wrapper for the ListFsuDiscoveryTargets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuDiscoveryTargets.go.html to see an example of how to use ListFsuDiscoveryTargetsRequest.
type ListFsuDiscoveryTargetsRequest struct {

	// Unique Exadata Fleet Update Discovery identifier.
	FsuDiscoveryId *string `mandatory:"true" contributesTo:"path" name:"fsuDiscoveryId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return a resource whose target OCID matches the given OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only entries whose status matches the given status.
	Status ListFsuDiscoveryTargetsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuDiscoveryTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListFsuDiscoveryTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuDiscoveryTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuDiscoveryTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuDiscoveryTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuDiscoveryTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuDiscoveryTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuDiscoveryTargetsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListFsuDiscoveryTargetsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuDiscoveryTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuDiscoveryTargetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuDiscoveryTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuDiscoveryTargetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuDiscoveryTargetsResponse wrapper for the ListFsuDiscoveryTargets operation
type ListFsuDiscoveryTargetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of TargetSummaryCollection instances
	TargetSummaryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListFsuDiscoveryTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuDiscoveryTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuDiscoveryTargetsStatusEnum Enum with underlying type: string
type ListFsuDiscoveryTargetsStatusEnum string

// Set of constants representing the allowable values for ListFsuDiscoveryTargetsStatusEnum
const (
	ListFsuDiscoveryTargetsStatusIdle         ListFsuDiscoveryTargetsStatusEnum = "IDLE"
	ListFsuDiscoveryTargetsStatusExecutingJob ListFsuDiscoveryTargetsStatusEnum = "EXECUTING_JOB"
	ListFsuDiscoveryTargetsStatusJobFailed    ListFsuDiscoveryTargetsStatusEnum = "JOB_FAILED"
)

var mappingListFsuDiscoveryTargetsStatusEnum = map[string]ListFsuDiscoveryTargetsStatusEnum{
	"IDLE":          ListFsuDiscoveryTargetsStatusIdle,
	"EXECUTING_JOB": ListFsuDiscoveryTargetsStatusExecutingJob,
	"JOB_FAILED":    ListFsuDiscoveryTargetsStatusJobFailed,
}

var mappingListFsuDiscoveryTargetsStatusEnumLowerCase = map[string]ListFsuDiscoveryTargetsStatusEnum{
	"idle":          ListFsuDiscoveryTargetsStatusIdle,
	"executing_job": ListFsuDiscoveryTargetsStatusExecutingJob,
	"job_failed":    ListFsuDiscoveryTargetsStatusJobFailed,
}

// GetListFsuDiscoveryTargetsStatusEnumValues Enumerates the set of values for ListFsuDiscoveryTargetsStatusEnum
func GetListFsuDiscoveryTargetsStatusEnumValues() []ListFsuDiscoveryTargetsStatusEnum {
	values := make([]ListFsuDiscoveryTargetsStatusEnum, 0)
	for _, v := range mappingListFsuDiscoveryTargetsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuDiscoveryTargetsStatusEnumStringValues Enumerates the set of values in String for ListFsuDiscoveryTargetsStatusEnum
func GetListFsuDiscoveryTargetsStatusEnumStringValues() []string {
	return []string{
		"IDLE",
		"EXECUTING_JOB",
		"JOB_FAILED",
	}
}

// GetMappingListFsuDiscoveryTargetsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuDiscoveryTargetsStatusEnum(val string) (ListFsuDiscoveryTargetsStatusEnum, bool) {
	enum, ok := mappingListFsuDiscoveryTargetsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuDiscoveryTargetsSortOrderEnum Enum with underlying type: string
type ListFsuDiscoveryTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListFsuDiscoveryTargetsSortOrderEnum
const (
	ListFsuDiscoveryTargetsSortOrderAsc  ListFsuDiscoveryTargetsSortOrderEnum = "ASC"
	ListFsuDiscoveryTargetsSortOrderDesc ListFsuDiscoveryTargetsSortOrderEnum = "DESC"
)

var mappingListFsuDiscoveryTargetsSortOrderEnum = map[string]ListFsuDiscoveryTargetsSortOrderEnum{
	"ASC":  ListFsuDiscoveryTargetsSortOrderAsc,
	"DESC": ListFsuDiscoveryTargetsSortOrderDesc,
}

var mappingListFsuDiscoveryTargetsSortOrderEnumLowerCase = map[string]ListFsuDiscoveryTargetsSortOrderEnum{
	"asc":  ListFsuDiscoveryTargetsSortOrderAsc,
	"desc": ListFsuDiscoveryTargetsSortOrderDesc,
}

// GetListFsuDiscoveryTargetsSortOrderEnumValues Enumerates the set of values for ListFsuDiscoveryTargetsSortOrderEnum
func GetListFsuDiscoveryTargetsSortOrderEnumValues() []ListFsuDiscoveryTargetsSortOrderEnum {
	values := make([]ListFsuDiscoveryTargetsSortOrderEnum, 0)
	for _, v := range mappingListFsuDiscoveryTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuDiscoveryTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListFsuDiscoveryTargetsSortOrderEnum
func GetListFsuDiscoveryTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuDiscoveryTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuDiscoveryTargetsSortOrderEnum(val string) (ListFsuDiscoveryTargetsSortOrderEnum, bool) {
	enum, ok := mappingListFsuDiscoveryTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuDiscoveryTargetsSortByEnum Enum with underlying type: string
type ListFsuDiscoveryTargetsSortByEnum string

// Set of constants representing the allowable values for ListFsuDiscoveryTargetsSortByEnum
const (
	ListFsuDiscoveryTargetsSortByCurrentversion ListFsuDiscoveryTargetsSortByEnum = "currentVersion"
	ListFsuDiscoveryTargetsSortByStatus         ListFsuDiscoveryTargetsSortByEnum = "status"
)

var mappingListFsuDiscoveryTargetsSortByEnum = map[string]ListFsuDiscoveryTargetsSortByEnum{
	"currentVersion": ListFsuDiscoveryTargetsSortByCurrentversion,
	"status":         ListFsuDiscoveryTargetsSortByStatus,
}

var mappingListFsuDiscoveryTargetsSortByEnumLowerCase = map[string]ListFsuDiscoveryTargetsSortByEnum{
	"currentversion": ListFsuDiscoveryTargetsSortByCurrentversion,
	"status":         ListFsuDiscoveryTargetsSortByStatus,
}

// GetListFsuDiscoveryTargetsSortByEnumValues Enumerates the set of values for ListFsuDiscoveryTargetsSortByEnum
func GetListFsuDiscoveryTargetsSortByEnumValues() []ListFsuDiscoveryTargetsSortByEnum {
	values := make([]ListFsuDiscoveryTargetsSortByEnum, 0)
	for _, v := range mappingListFsuDiscoveryTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuDiscoveryTargetsSortByEnumStringValues Enumerates the set of values in String for ListFsuDiscoveryTargetsSortByEnum
func GetListFsuDiscoveryTargetsSortByEnumStringValues() []string {
	return []string{
		"currentVersion",
		"status",
	}
}

// GetMappingListFsuDiscoveryTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuDiscoveryTargetsSortByEnum(val string) (ListFsuDiscoveryTargetsSortByEnum, bool) {
	enum, ok := mappingListFsuDiscoveryTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
