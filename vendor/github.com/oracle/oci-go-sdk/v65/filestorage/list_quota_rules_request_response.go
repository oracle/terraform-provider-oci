// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListQuotaRulesRequest wrapper for the ListQuotaRules operation
//
// # See also
//
// Click https://docs.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/filestorage/ListQuotaRules.go.html to see an example of how to use ListQuotaRulesRequest.
type ListQuotaRulesRequest struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`

	// The type of the owner of this quota rule and usage.
	PrincipalType ListQuotaRulesPrincipalTypeEnum `mandatory:"true" contributesTo:"query" name:"principalType" omitEmpty:"true"`

	// For list pagination. The maximum number of results per page,
	// or items to return in a paginated "List" call.
	// 1 is the minimum, 4096 is the maximum.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The value of the `opc-next-page` response
	// header from the previous "List" call.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// An identifier for the user or the group associated with quota rule and usage. UNIX-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"false" contributesTo:"query" name:"principalId"`

	// An option to display only the users or groups that violate their quota rules.
	// If `areViolatorsOnly` is false, results report all the quota and usage.
	// If `areViolatorsOnly` is true, results only report the quota and usage for
	// the users or groups that violate their quota rules.
	AreViolatorsOnly *bool `mandatory:"false" contributesTo:"query" name:"areViolatorsOnly"`

	// The sort order to use, either 'asc' or 'desc', where 'asc' is
	// ascending and 'desc' is descending. The default order is 'desc'
	// except for numeric values.
	SortOrder ListQuotaRulesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// For optimistic concurrency control. In the PUT or DELETE call
	// for a resource, set the `if-match` parameter to the value of the
	// etag from a previous GET or POST response for that resource.
	// The resource will be updated or deleted only if the etag you
	// provide matches the resource's current etag value.
	IfMatch *string `mandatory:"false" contributesTo:"header" name:"if-match"`

	// Unique identifier for the request.
	// If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListQuotaRulesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListQuotaRulesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListQuotaRulesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListQuotaRulesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListQuotaRulesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListQuotaRulesPrincipalTypeEnum(string(request.PrincipalType)); !ok && request.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", request.PrincipalType, strings.Join(GetListQuotaRulesPrincipalTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListQuotaRulesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListQuotaRulesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListQuotaRulesResponse wrapper for the ListQuotaRules operation
type ListQuotaRulesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []QuotaRuleSummary instances
	Items []QuotaRuleSummary `presentIn:"body"`

	// For list pagination. When this header appears in the response,
	// additional pages of results remain.
	// For important details about how pagination works,
	// see List Pagination (https://docs.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListQuotaRulesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListQuotaRulesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListQuotaRulesPrincipalTypeEnum Enum with underlying type: string
type ListQuotaRulesPrincipalTypeEnum string

// Set of constants representing the allowable values for ListQuotaRulesPrincipalTypeEnum
const (
	ListQuotaRulesPrincipalTypeFileSystemLevel ListQuotaRulesPrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	ListQuotaRulesPrincipalTypeDefaultGroup    ListQuotaRulesPrincipalTypeEnum = "DEFAULT_GROUP"
	ListQuotaRulesPrincipalTypeDefaultUser     ListQuotaRulesPrincipalTypeEnum = "DEFAULT_USER"
	ListQuotaRulesPrincipalTypeIndividualGroup ListQuotaRulesPrincipalTypeEnum = "INDIVIDUAL_GROUP"
	ListQuotaRulesPrincipalTypeIndividualUser  ListQuotaRulesPrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingListQuotaRulesPrincipalTypeEnum = map[string]ListQuotaRulesPrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": ListQuotaRulesPrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     ListQuotaRulesPrincipalTypeDefaultGroup,
	"DEFAULT_USER":      ListQuotaRulesPrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  ListQuotaRulesPrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   ListQuotaRulesPrincipalTypeIndividualUser,
}

var mappingListQuotaRulesPrincipalTypeEnumLowerCase = map[string]ListQuotaRulesPrincipalTypeEnum{
	"file_system_level": ListQuotaRulesPrincipalTypeFileSystemLevel,
	"default_group":     ListQuotaRulesPrincipalTypeDefaultGroup,
	"default_user":      ListQuotaRulesPrincipalTypeDefaultUser,
	"individual_group":  ListQuotaRulesPrincipalTypeIndividualGroup,
	"individual_user":   ListQuotaRulesPrincipalTypeIndividualUser,
}

// GetListQuotaRulesPrincipalTypeEnumValues Enumerates the set of values for ListQuotaRulesPrincipalTypeEnum
func GetListQuotaRulesPrincipalTypeEnumValues() []ListQuotaRulesPrincipalTypeEnum {
	values := make([]ListQuotaRulesPrincipalTypeEnum, 0)
	for _, v := range mappingListQuotaRulesPrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetListQuotaRulesPrincipalTypeEnumStringValues Enumerates the set of values in String for ListQuotaRulesPrincipalTypeEnum
func GetListQuotaRulesPrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingListQuotaRulesPrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQuotaRulesPrincipalTypeEnum(val string) (ListQuotaRulesPrincipalTypeEnum, bool) {
	enum, ok := mappingListQuotaRulesPrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListQuotaRulesSortOrderEnum Enum with underlying type: string
type ListQuotaRulesSortOrderEnum string

// Set of constants representing the allowable values for ListQuotaRulesSortOrderEnum
const (
	ListQuotaRulesSortOrderAsc  ListQuotaRulesSortOrderEnum = "ASC"
	ListQuotaRulesSortOrderDesc ListQuotaRulesSortOrderEnum = "DESC"
)

var mappingListQuotaRulesSortOrderEnum = map[string]ListQuotaRulesSortOrderEnum{
	"ASC":  ListQuotaRulesSortOrderAsc,
	"DESC": ListQuotaRulesSortOrderDesc,
}

var mappingListQuotaRulesSortOrderEnumLowerCase = map[string]ListQuotaRulesSortOrderEnum{
	"asc":  ListQuotaRulesSortOrderAsc,
	"desc": ListQuotaRulesSortOrderDesc,
}

// GetListQuotaRulesSortOrderEnumValues Enumerates the set of values for ListQuotaRulesSortOrderEnum
func GetListQuotaRulesSortOrderEnumValues() []ListQuotaRulesSortOrderEnum {
	values := make([]ListQuotaRulesSortOrderEnum, 0)
	for _, v := range mappingListQuotaRulesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListQuotaRulesSortOrderEnumStringValues Enumerates the set of values in String for ListQuotaRulesSortOrderEnum
func GetListQuotaRulesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListQuotaRulesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListQuotaRulesSortOrderEnum(val string) (ListQuotaRulesSortOrderEnum, bool) {
	enum, ok := mappingListQuotaRulesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
