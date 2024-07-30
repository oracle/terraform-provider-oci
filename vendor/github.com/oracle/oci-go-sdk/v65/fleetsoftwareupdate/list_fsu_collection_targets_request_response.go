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

// ListFsuCollectionTargetsRequest wrapper for the ListFsuCollectionTargets operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetsoftwareupdate/ListFsuCollectionTargets.go.html to see an example of how to use ListFsuCollectionTargetsRequest.
type ListFsuCollectionTargetsRequest struct {

	// Unique Exadata Fleet Update Collection identifier.
	FsuCollectionId *string `mandatory:"true" contributesTo:"path" name:"fsuCollectionId"`

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return a resource whose target OCID matches the given OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only entries whose status matches the given status.
	Status ListFsuCollectionTargetsStatusEnum `mandatory:"false" contributesTo:"query" name:"status" omitEmpty:"true"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results.
	// This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListFsuCollectionTargetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided.
	SortBy ListFsuCollectionTargetsSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFsuCollectionTargetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFsuCollectionTargetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFsuCollectionTargetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFsuCollectionTargetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFsuCollectionTargetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFsuCollectionTargetsStatusEnum(string(request.Status)); !ok && request.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", request.Status, strings.Join(GetListFsuCollectionTargetsStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCollectionTargetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListFsuCollectionTargetsSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFsuCollectionTargetsSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListFsuCollectionTargetsSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFsuCollectionTargetsResponse wrapper for the ListFsuCollectionTargets operation
type ListFsuCollectionTargetsResponse struct {

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

func (response ListFsuCollectionTargetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFsuCollectionTargetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFsuCollectionTargetsStatusEnum Enum with underlying type: string
type ListFsuCollectionTargetsStatusEnum string

// Set of constants representing the allowable values for ListFsuCollectionTargetsStatusEnum
const (
	ListFsuCollectionTargetsStatusIdle         ListFsuCollectionTargetsStatusEnum = "IDLE"
	ListFsuCollectionTargetsStatusExecutingJob ListFsuCollectionTargetsStatusEnum = "EXECUTING_JOB"
	ListFsuCollectionTargetsStatusJobFailed    ListFsuCollectionTargetsStatusEnum = "JOB_FAILED"
)

var mappingListFsuCollectionTargetsStatusEnum = map[string]ListFsuCollectionTargetsStatusEnum{
	"IDLE":          ListFsuCollectionTargetsStatusIdle,
	"EXECUTING_JOB": ListFsuCollectionTargetsStatusExecutingJob,
	"JOB_FAILED":    ListFsuCollectionTargetsStatusJobFailed,
}

var mappingListFsuCollectionTargetsStatusEnumLowerCase = map[string]ListFsuCollectionTargetsStatusEnum{
	"idle":          ListFsuCollectionTargetsStatusIdle,
	"executing_job": ListFsuCollectionTargetsStatusExecutingJob,
	"job_failed":    ListFsuCollectionTargetsStatusJobFailed,
}

// GetListFsuCollectionTargetsStatusEnumValues Enumerates the set of values for ListFsuCollectionTargetsStatusEnum
func GetListFsuCollectionTargetsStatusEnumValues() []ListFsuCollectionTargetsStatusEnum {
	values := make([]ListFsuCollectionTargetsStatusEnum, 0)
	for _, v := range mappingListFsuCollectionTargetsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCollectionTargetsStatusEnumStringValues Enumerates the set of values in String for ListFsuCollectionTargetsStatusEnum
func GetListFsuCollectionTargetsStatusEnumStringValues() []string {
	return []string{
		"IDLE",
		"EXECUTING_JOB",
		"JOB_FAILED",
	}
}

// GetMappingListFsuCollectionTargetsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCollectionTargetsStatusEnum(val string) (ListFsuCollectionTargetsStatusEnum, bool) {
	enum, ok := mappingListFsuCollectionTargetsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCollectionTargetsSortOrderEnum Enum with underlying type: string
type ListFsuCollectionTargetsSortOrderEnum string

// Set of constants representing the allowable values for ListFsuCollectionTargetsSortOrderEnum
const (
	ListFsuCollectionTargetsSortOrderAsc  ListFsuCollectionTargetsSortOrderEnum = "ASC"
	ListFsuCollectionTargetsSortOrderDesc ListFsuCollectionTargetsSortOrderEnum = "DESC"
)

var mappingListFsuCollectionTargetsSortOrderEnum = map[string]ListFsuCollectionTargetsSortOrderEnum{
	"ASC":  ListFsuCollectionTargetsSortOrderAsc,
	"DESC": ListFsuCollectionTargetsSortOrderDesc,
}

var mappingListFsuCollectionTargetsSortOrderEnumLowerCase = map[string]ListFsuCollectionTargetsSortOrderEnum{
	"asc":  ListFsuCollectionTargetsSortOrderAsc,
	"desc": ListFsuCollectionTargetsSortOrderDesc,
}

// GetListFsuCollectionTargetsSortOrderEnumValues Enumerates the set of values for ListFsuCollectionTargetsSortOrderEnum
func GetListFsuCollectionTargetsSortOrderEnumValues() []ListFsuCollectionTargetsSortOrderEnum {
	values := make([]ListFsuCollectionTargetsSortOrderEnum, 0)
	for _, v := range mappingListFsuCollectionTargetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCollectionTargetsSortOrderEnumStringValues Enumerates the set of values in String for ListFsuCollectionTargetsSortOrderEnum
func GetListFsuCollectionTargetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListFsuCollectionTargetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCollectionTargetsSortOrderEnum(val string) (ListFsuCollectionTargetsSortOrderEnum, bool) {
	enum, ok := mappingListFsuCollectionTargetsSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFsuCollectionTargetsSortByEnum Enum with underlying type: string
type ListFsuCollectionTargetsSortByEnum string

// Set of constants representing the allowable values for ListFsuCollectionTargetsSortByEnum
const (
	ListFsuCollectionTargetsSortByCurrentversion ListFsuCollectionTargetsSortByEnum = "currentVersion"
	ListFsuCollectionTargetsSortByStatus         ListFsuCollectionTargetsSortByEnum = "status"
)

var mappingListFsuCollectionTargetsSortByEnum = map[string]ListFsuCollectionTargetsSortByEnum{
	"currentVersion": ListFsuCollectionTargetsSortByCurrentversion,
	"status":         ListFsuCollectionTargetsSortByStatus,
}

var mappingListFsuCollectionTargetsSortByEnumLowerCase = map[string]ListFsuCollectionTargetsSortByEnum{
	"currentversion": ListFsuCollectionTargetsSortByCurrentversion,
	"status":         ListFsuCollectionTargetsSortByStatus,
}

// GetListFsuCollectionTargetsSortByEnumValues Enumerates the set of values for ListFsuCollectionTargetsSortByEnum
func GetListFsuCollectionTargetsSortByEnumValues() []ListFsuCollectionTargetsSortByEnum {
	values := make([]ListFsuCollectionTargetsSortByEnum, 0)
	for _, v := range mappingListFsuCollectionTargetsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListFsuCollectionTargetsSortByEnumStringValues Enumerates the set of values in String for ListFsuCollectionTargetsSortByEnum
func GetListFsuCollectionTargetsSortByEnumStringValues() []string {
	return []string{
		"currentVersion",
		"status",
	}
}

// GetMappingListFsuCollectionTargetsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFsuCollectionTargetsSortByEnum(val string) (ListFsuCollectionTargetsSortByEnum, bool) {
	enum, ok := mappingListFsuCollectionTargetsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
