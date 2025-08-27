// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListVulnerabilitiesRequest wrapper for the ListVulnerabilities operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/dblm/ListVulnerabilities.go.html to see an example of how to use ListVulnerabilitiesRequest.
type ListVulnerabilitiesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListVulnerabilitiesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for vulnerableResources is descending. Default order for cveId is descending.
	SortBy ListVulnerabilitiesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// A token that uniquely identifies a request so it can be retried in case of a timeout or
	// server error without risk of executing that same action again. Retry tokens expire after 24
	// hours, but can be invalidated before then due to conflicting operations. For example, if a resource
	// has been deleted and purged from the system, then a retry of the original creation request
	// might be rejected.
	OpcRetryToken *string `mandatory:"false" contributesTo:"header" name:"opc-retry-token"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only resources their lifecycleState matches the given lifecycleState.
	LifecycleState DblmVulnerabilityLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only database that match the given release version.
	DatabaseRelease *string `mandatory:"false" contributesTo:"query" name:"databaseRelease"`

	// A filter to return only resources that match the given resource id.
	ResourceId *string `mandatory:"false" contributesTo:"query" name:"resourceId"`

	// Filter by one or more severity types.
	// Possible values are critical, high, medium, low, info.
	SeverityType []VulnerabilityRiskLevelEnum `contributesTo:"query" name:"severityType" omitEmpty:"true" collectionFormat:"multi"`

	// The search input for filter cve id and cve description.
	SearchBy *string `mandatory:"false" contributesTo:"query" name:"searchBy"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListVulnerabilitiesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListVulnerabilitiesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListVulnerabilitiesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListVulnerabilitiesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListVulnerabilitiesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListVulnerabilitiesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListVulnerabilitiesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListVulnerabilitiesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListVulnerabilitiesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDblmVulnerabilityLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetDblmVulnerabilityLifecycleStateEnumStringValues(), ",")))
	}
	for _, val := range request.SeverityType {
		if _, ok := GetMappingVulnerabilityRiskLevelEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SeverityType: %s. Supported values are: %s.", val, strings.Join(GetVulnerabilityRiskLevelEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListVulnerabilitiesResponse wrapper for the ListVulnerabilities operation
type ListVulnerabilitiesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of VulnerabilityCollection instances
	VulnerabilityCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListVulnerabilitiesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListVulnerabilitiesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListVulnerabilitiesSortOrderEnum Enum with underlying type: string
type ListVulnerabilitiesSortOrderEnum string

// Set of constants representing the allowable values for ListVulnerabilitiesSortOrderEnum
const (
	ListVulnerabilitiesSortOrderAsc  ListVulnerabilitiesSortOrderEnum = "ASC"
	ListVulnerabilitiesSortOrderDesc ListVulnerabilitiesSortOrderEnum = "DESC"
)

var mappingListVulnerabilitiesSortOrderEnum = map[string]ListVulnerabilitiesSortOrderEnum{
	"ASC":  ListVulnerabilitiesSortOrderAsc,
	"DESC": ListVulnerabilitiesSortOrderDesc,
}

var mappingListVulnerabilitiesSortOrderEnumLowerCase = map[string]ListVulnerabilitiesSortOrderEnum{
	"asc":  ListVulnerabilitiesSortOrderAsc,
	"desc": ListVulnerabilitiesSortOrderDesc,
}

// GetListVulnerabilitiesSortOrderEnumValues Enumerates the set of values for ListVulnerabilitiesSortOrderEnum
func GetListVulnerabilitiesSortOrderEnumValues() []ListVulnerabilitiesSortOrderEnum {
	values := make([]ListVulnerabilitiesSortOrderEnum, 0)
	for _, v := range mappingListVulnerabilitiesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListVulnerabilitiesSortOrderEnumStringValues Enumerates the set of values in String for ListVulnerabilitiesSortOrderEnum
func GetListVulnerabilitiesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListVulnerabilitiesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVulnerabilitiesSortOrderEnum(val string) (ListVulnerabilitiesSortOrderEnum, bool) {
	enum, ok := mappingListVulnerabilitiesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListVulnerabilitiesSortByEnum Enum with underlying type: string
type ListVulnerabilitiesSortByEnum string

// Set of constants representing the allowable values for ListVulnerabilitiesSortByEnum
const (
	ListVulnerabilitiesSortByCveid               ListVulnerabilitiesSortByEnum = "cveId"
	ListVulnerabilitiesSortByVulnerableresources ListVulnerabilitiesSortByEnum = "vulnerableResources"
)

var mappingListVulnerabilitiesSortByEnum = map[string]ListVulnerabilitiesSortByEnum{
	"cveId":               ListVulnerabilitiesSortByCveid,
	"vulnerableResources": ListVulnerabilitiesSortByVulnerableresources,
}

var mappingListVulnerabilitiesSortByEnumLowerCase = map[string]ListVulnerabilitiesSortByEnum{
	"cveid":               ListVulnerabilitiesSortByCveid,
	"vulnerableresources": ListVulnerabilitiesSortByVulnerableresources,
}

// GetListVulnerabilitiesSortByEnumValues Enumerates the set of values for ListVulnerabilitiesSortByEnum
func GetListVulnerabilitiesSortByEnumValues() []ListVulnerabilitiesSortByEnum {
	values := make([]ListVulnerabilitiesSortByEnum, 0)
	for _, v := range mappingListVulnerabilitiesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListVulnerabilitiesSortByEnumStringValues Enumerates the set of values in String for ListVulnerabilitiesSortByEnum
func GetListVulnerabilitiesSortByEnumStringValues() []string {
	return []string{
		"cveId",
		"vulnerableResources",
	}
}

// GetMappingListVulnerabilitiesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListVulnerabilitiesSortByEnum(val string) (ListVulnerabilitiesSortByEnum, bool) {
	enum, ok := mappingListVulnerabilitiesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
