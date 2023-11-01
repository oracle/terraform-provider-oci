// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListProfileSummariesRequest wrapper for the ListProfileSummaries operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListProfileSummaries.go.html to see an example of how to use ListProfileSummariesRequest.
type ListProfileSummariesRequest struct {

	// The OCID of the user assessment.
	UserAssessmentId *string `mandatory:"true" contributesTo:"path" name:"userAssessmentId"`

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListProfileSummariesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// A filter to return only items that match the specified profile name.
	ProfileName *string `mandatory:"false" contributesTo:"query" name:"profileName"`

	// An optional filter to return the user created profiles.
	IsUserCreated *bool `mandatory:"false" contributesTo:"query" name:"isUserCreated"`

	// An optional filter to filter the profiles based on password verification function.
	PasswordVerificationFunction *string `mandatory:"false" contributesTo:"query" name:"passwordVerificationFunction"`

	// An optional filter to return the profiles having user count greater than or equal to the provided value.
	UserCountGreaterThanOrEqual *string `mandatory:"false" contributesTo:"query" name:"userCountGreaterThanOrEqual"`

	// An optional filter to return the profiles having user count less than the provided value.
	UserCountLessThan *string `mandatory:"false" contributesTo:"query" name:"userCountLessThan"`

	// An optional filter to return the profiles having allow failed login attempts number greater than or equal to the provided value.
	// String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	FailedLoginAttemptsGreaterThanOrEqual *string `mandatory:"false" contributesTo:"query" name:"failedLoginAttemptsGreaterThanOrEqual"`

	// An optional filter to return the profiles having failed login attempts number less than the provided value.
	// String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	FailedLoginAttemptsLessThan *string `mandatory:"false" contributesTo:"query" name:"failedLoginAttemptsLessThan"`

	// An optional filter to return the profiles permitting the user to spawn multiple sessions having count.
	// greater than or equal to the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	SessionsPerUserGreaterThanOrEqual *string `mandatory:"false" contributesTo:"query" name:"sessionsPerUserGreaterThanOrEqual"`

	// An optional filter to return the profiles permitting the user to spawn multiple sessions having count less than
	// the provided value. String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	SessionsPerUserLessThan *string `mandatory:"false" contributesTo:"query" name:"sessionsPerUserLessThan"`

	// An optional filter to return the profiles allowing inactive account time in days greater than or equal to the provided value.
	// String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	InactiveAccountTimeGreaterThanOrEqual *string `mandatory:"false" contributesTo:"query" name:"inactiveAccountTimeGreaterThanOrEqual"`

	// An optional filter to return the profiles  allowing inactive account time in days less than the provided value.
	// String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	InactiveAccountTimeLessThan *string `mandatory:"false" contributesTo:"query" name:"inactiveAccountTimeLessThan"`

	// An optional filter to return the profiles having password lock number greater than or equal to the provided value.
	// String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	PasswordLockTimeGreaterThanOrEqual *string `mandatory:"false" contributesTo:"query" name:"passwordLockTimeGreaterThanOrEqual"`

	// An optional filter to return the profiles having password lock number less than the provided value.
	// String value is used for accommodating the "UNLIMITED" and "DEFAULT" values.
	PasswordLockTimeLessThan *string `mandatory:"false" contributesTo:"query" name:"passwordLockTimeLessThan"`

	// The field to sort by. You can specify only one sort order (sortOrder). The default order is targetId ASC.
	SortBy ListProfileSummariesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// The sort order to use, either ascending (ASC) or descending (DESC).
	SortOrder ListProfileSummariesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListProfileSummariesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListProfileSummariesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListProfileSummariesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListProfileSummariesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListProfileSummariesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListProfileSummariesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListProfileSummariesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfileSummariesSortByEnum(string(request.SortBy)); !ok && request.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", request.SortBy, strings.Join(GetListProfileSummariesSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListProfileSummariesSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListProfileSummariesSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListProfileSummariesResponse wrapper for the ListProfileSummaries operation
type ListProfileSummariesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []ProfileSummary instances
	Items []ProfileSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListProfileSummariesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListProfileSummariesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListProfileSummariesAccessLevelEnum Enum with underlying type: string
type ListProfileSummariesAccessLevelEnum string

// Set of constants representing the allowable values for ListProfileSummariesAccessLevelEnum
const (
	ListProfileSummariesAccessLevelRestricted ListProfileSummariesAccessLevelEnum = "RESTRICTED"
	ListProfileSummariesAccessLevelAccessible ListProfileSummariesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListProfileSummariesAccessLevelEnum = map[string]ListProfileSummariesAccessLevelEnum{
	"RESTRICTED": ListProfileSummariesAccessLevelRestricted,
	"ACCESSIBLE": ListProfileSummariesAccessLevelAccessible,
}

var mappingListProfileSummariesAccessLevelEnumLowerCase = map[string]ListProfileSummariesAccessLevelEnum{
	"restricted": ListProfileSummariesAccessLevelRestricted,
	"accessible": ListProfileSummariesAccessLevelAccessible,
}

// GetListProfileSummariesAccessLevelEnumValues Enumerates the set of values for ListProfileSummariesAccessLevelEnum
func GetListProfileSummariesAccessLevelEnumValues() []ListProfileSummariesAccessLevelEnum {
	values := make([]ListProfileSummariesAccessLevelEnum, 0)
	for _, v := range mappingListProfileSummariesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileSummariesAccessLevelEnumStringValues Enumerates the set of values in String for ListProfileSummariesAccessLevelEnum
func GetListProfileSummariesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListProfileSummariesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileSummariesAccessLevelEnum(val string) (ListProfileSummariesAccessLevelEnum, bool) {
	enum, ok := mappingListProfileSummariesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfileSummariesSortByEnum Enum with underlying type: string
type ListProfileSummariesSortByEnum string

// Set of constants representing the allowable values for ListProfileSummariesSortByEnum
const (
	ListProfileSummariesSortByProfilename                  ListProfileSummariesSortByEnum = "profileName"
	ListProfileSummariesSortByTargetid                     ListProfileSummariesSortByEnum = "targetId"
	ListProfileSummariesSortByIsusercreated                ListProfileSummariesSortByEnum = "isUserCreated"
	ListProfileSummariesSortByPasswordverificationfunction ListProfileSummariesSortByEnum = "passwordVerificationFunction"
	ListProfileSummariesSortByUsercount                    ListProfileSummariesSortByEnum = "userCount"
	ListProfileSummariesSortBySessionsperuser              ListProfileSummariesSortByEnum = "sessionsPerUser"
	ListProfileSummariesSortByInactiveaccounttime          ListProfileSummariesSortByEnum = "inactiveAccountTime"
	ListProfileSummariesSortByPasswordlocktime             ListProfileSummariesSortByEnum = "passwordLockTime"
	ListProfileSummariesSortByFailedloginattempts          ListProfileSummariesSortByEnum = "failedLoginAttempts"
)

var mappingListProfileSummariesSortByEnum = map[string]ListProfileSummariesSortByEnum{
	"profileName":                  ListProfileSummariesSortByProfilename,
	"targetId":                     ListProfileSummariesSortByTargetid,
	"isUserCreated":                ListProfileSummariesSortByIsusercreated,
	"passwordVerificationFunction": ListProfileSummariesSortByPasswordverificationfunction,
	"userCount":                    ListProfileSummariesSortByUsercount,
	"sessionsPerUser":              ListProfileSummariesSortBySessionsperuser,
	"inactiveAccountTime":          ListProfileSummariesSortByInactiveaccounttime,
	"passwordLockTime":             ListProfileSummariesSortByPasswordlocktime,
	"failedLoginAttempts":          ListProfileSummariesSortByFailedloginattempts,
}

var mappingListProfileSummariesSortByEnumLowerCase = map[string]ListProfileSummariesSortByEnum{
	"profilename":                  ListProfileSummariesSortByProfilename,
	"targetid":                     ListProfileSummariesSortByTargetid,
	"isusercreated":                ListProfileSummariesSortByIsusercreated,
	"passwordverificationfunction": ListProfileSummariesSortByPasswordverificationfunction,
	"usercount":                    ListProfileSummariesSortByUsercount,
	"sessionsperuser":              ListProfileSummariesSortBySessionsperuser,
	"inactiveaccounttime":          ListProfileSummariesSortByInactiveaccounttime,
	"passwordlocktime":             ListProfileSummariesSortByPasswordlocktime,
	"failedloginattempts":          ListProfileSummariesSortByFailedloginattempts,
}

// GetListProfileSummariesSortByEnumValues Enumerates the set of values for ListProfileSummariesSortByEnum
func GetListProfileSummariesSortByEnumValues() []ListProfileSummariesSortByEnum {
	values := make([]ListProfileSummariesSortByEnum, 0)
	for _, v := range mappingListProfileSummariesSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileSummariesSortByEnumStringValues Enumerates the set of values in String for ListProfileSummariesSortByEnum
func GetListProfileSummariesSortByEnumStringValues() []string {
	return []string{
		"profileName",
		"targetId",
		"isUserCreated",
		"passwordVerificationFunction",
		"userCount",
		"sessionsPerUser",
		"inactiveAccountTime",
		"passwordLockTime",
		"failedLoginAttempts",
	}
}

// GetMappingListProfileSummariesSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileSummariesSortByEnum(val string) (ListProfileSummariesSortByEnum, bool) {
	enum, ok := mappingListProfileSummariesSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListProfileSummariesSortOrderEnum Enum with underlying type: string
type ListProfileSummariesSortOrderEnum string

// Set of constants representing the allowable values for ListProfileSummariesSortOrderEnum
const (
	ListProfileSummariesSortOrderAsc  ListProfileSummariesSortOrderEnum = "ASC"
	ListProfileSummariesSortOrderDesc ListProfileSummariesSortOrderEnum = "DESC"
)

var mappingListProfileSummariesSortOrderEnum = map[string]ListProfileSummariesSortOrderEnum{
	"ASC":  ListProfileSummariesSortOrderAsc,
	"DESC": ListProfileSummariesSortOrderDesc,
}

var mappingListProfileSummariesSortOrderEnumLowerCase = map[string]ListProfileSummariesSortOrderEnum{
	"asc":  ListProfileSummariesSortOrderAsc,
	"desc": ListProfileSummariesSortOrderDesc,
}

// GetListProfileSummariesSortOrderEnumValues Enumerates the set of values for ListProfileSummariesSortOrderEnum
func GetListProfileSummariesSortOrderEnumValues() []ListProfileSummariesSortOrderEnum {
	values := make([]ListProfileSummariesSortOrderEnum, 0)
	for _, v := range mappingListProfileSummariesSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListProfileSummariesSortOrderEnumStringValues Enumerates the set of values in String for ListProfileSummariesSortOrderEnum
func GetListProfileSummariesSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListProfileSummariesSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListProfileSummariesSortOrderEnum(val string) (ListProfileSummariesSortOrderEnum, bool) {
	enum, ok := mappingListProfileSummariesSortOrderEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
