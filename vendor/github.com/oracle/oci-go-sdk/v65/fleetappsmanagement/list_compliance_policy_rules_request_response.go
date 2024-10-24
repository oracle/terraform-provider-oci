// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListCompliancePolicyRulesRequest wrapper for the ListCompliancePolicyRules operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/fleetappsmanagement/ListCompliancePolicyRules.go.html to see an example of how to use ListCompliancePolicyRulesRequest.
type ListCompliancePolicyRulesRequest struct {

	// The ID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// A filter to return only resources that match the entire display name given.
	DisplayName *string `mandatory:"false" contributesTo:"query" name:"displayName"`

	// A filter to return only resources that match the given lifecycle state. The
	// state value is case-insensitive.
	LifecycleState CompliancePolicyRuleLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the patch selection against the given patch name.
	PatchName *string `mandatory:"false" contributesTo:"query" name:"patchName"`

	// unique CompliancePolicy identifier.
	CompliancePolicyId *string `mandatory:"false" contributesTo:"query" name:"compliancePolicyId"`

	// unique CompliancePolicyRule identifier.
	Id *string `mandatory:"false" contributesTo:"query" name:"id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// A token representing the position at which to start retrieving results. This must come from the `opc-next-page` header field of a previous response.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder ListCompliancePolicyRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for timeCreated is descending. Default order for displayName is ascending.
	SortBy ListCompliancePolicyRulesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListCompliancePolicyRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListCompliancePolicyRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListCompliancePolicyRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListCompliancePolicyRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListCompliancePolicyRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCompliancePolicyRuleLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetCompliancePolicyRuleLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCompliancePolicyRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListCompliancePolicyRulesSortOrderEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListCompliancePolicyRulesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListCompliancePolicyRulesSortByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListCompliancePolicyRulesResponse wrapper for the ListCompliancePolicyRules operation
type ListCompliancePolicyRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of CompliancePolicyRuleCollection instances
	CompliancePolicyRuleCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact
	// Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListCompliancePolicyRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListCompliancePolicyRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListCompliancePolicyRulesSortOrderEnum Enum with underlying type: string
type ListCompliancePolicyRulesSortOrderEnum string

// Set of constants representing the allowable values for ListCompliancePolicyRulesSortOrderEnum
const (
	ListCompliancePolicyRulesSortOrderAsc  ListCompliancePolicyRulesSortOrderEnum = "ASC"
	ListCompliancePolicyRulesSortOrderDesc ListCompliancePolicyRulesSortOrderEnum = "DESC"
)

var mappingListCompliancePolicyRulesSortOrderEnum = map[string]ListCompliancePolicyRulesSortOrderEnum{
	"ASC":  ListCompliancePolicyRulesSortOrderAsc,
	"DESC": ListCompliancePolicyRulesSortOrderDesc,
}

var mappingListCompliancePolicyRulesSortOrderEnumLowerCase = map[string]ListCompliancePolicyRulesSortOrderEnum{
	"asc":  ListCompliancePolicyRulesSortOrderAsc,
	"desc": ListCompliancePolicyRulesSortOrderDesc,
}

// GetListCompliancePolicyRulesSortOrderEnumValues Enumerates the set of values for ListCompliancePolicyRulesSortOrderEnum
func GetListCompliancePolicyRulesSortOrderEnumValues() []ListCompliancePolicyRulesSortOrderEnum {
	values := make([]ListCompliancePolicyRulesSortOrderEnum, 0)
	for _, v := range mappingListCompliancePolicyRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListCompliancePolicyRulesSortOrderEnumStringValues Enumerates the set of values in String for ListCompliancePolicyRulesSortOrderEnum
func GetListCompliancePolicyRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListCompliancePolicyRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCompliancePolicyRulesSortOrderEnum(val string) (ListCompliancePolicyRulesSortOrderEnum, bool) {
	enum, ok := mappingListCompliancePolicyRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListCompliancePolicyRulesSortByEnum Enum with underlying type: string
type ListCompliancePolicyRulesSortByEnum string

// Set of constants representing the allowable values for ListCompliancePolicyRulesSortByEnum
const (
	ListCompliancePolicyRulesSortByTimecreated ListCompliancePolicyRulesSortByEnum = "timeCreated"
	ListCompliancePolicyRulesSortByDisplayname ListCompliancePolicyRulesSortByEnum = "displayName"
)

var mappingListCompliancePolicyRulesSortByEnum = map[string]ListCompliancePolicyRulesSortByEnum{
	"timeCreated": ListCompliancePolicyRulesSortByTimecreated,
	"displayName": ListCompliancePolicyRulesSortByDisplayname,
}

var mappingListCompliancePolicyRulesSortByEnumLowerCase = map[string]ListCompliancePolicyRulesSortByEnum{
	"timecreated": ListCompliancePolicyRulesSortByTimecreated,
	"displayname": ListCompliancePolicyRulesSortByDisplayname,
}

// GetListCompliancePolicyRulesSortByEnumValues Enumerates the set of values for ListCompliancePolicyRulesSortByEnum
func GetListCompliancePolicyRulesSortByEnumValues() []ListCompliancePolicyRulesSortByEnum {
	values := make([]ListCompliancePolicyRulesSortByEnum, 0)
	for _, v := range mappingListCompliancePolicyRulesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListCompliancePolicyRulesSortByEnumStringValues Enumerates the set of values in String for ListCompliancePolicyRulesSortByEnum
func GetListCompliancePolicyRulesSortByEnumStringValues() []string {
	return []string{
		"timeCreated",
		"displayName",
	}
}

// GetMappingListCompliancePolicyRulesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListCompliancePolicyRulesSortByEnum(val string) (ListCompliancePolicyRulesSortByEnum, bool) {
	enum, ok := mappingListCompliancePolicyRulesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
