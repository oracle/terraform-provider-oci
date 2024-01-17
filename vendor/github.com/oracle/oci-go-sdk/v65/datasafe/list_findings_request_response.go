// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ListFindingsRequest wrapper for the ListFindings operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListFindings.go.html to see an example of how to use ListFindingsRequest.
type ListFindingsRequest struct {

	// The OCID of the security assessment.
	SecurityAssessmentId *string `mandatory:"true" contributesTo:"path" name:"securityAssessmentId"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// A filter to return only the findings that are marked as top findings.
	IsTopFinding *bool `mandatory:"false" contributesTo:"query" name:"isTopFinding"`

	// A filter to return only findings of a particular risk level.
	Severity ListFindingsSeverityEnum `mandatory:"false" contributesTo:"query" name:"severity" omitEmpty:"true"`

	// A filter to return only the findings that match the specified lifecycle states.
	LifecycleState ListFindingsLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// An optional filter to return only findings that match the specified reference.
	References ListFindingsReferencesEnum `mandatory:"false" contributesTo:"query" name:"references" omitEmpty:"true"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListFindingsAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// Each finding in security assessment has an associated key (think of key as a finding's name).
	// For a given finding, the key will be the same across targets. The user can use these keys to filter the findings.
	FindingKey *string `mandatory:"false" contributesTo:"query" name:"findingKey"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListFindingsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListFindingsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListFindingsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListFindingsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListFindingsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListFindingsSeverityEnum(string(request.Severity)); !ok && request.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", request.Severity, strings.Join(GetListFindingsSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFindingsLifecycleStateEnum(string(request.LifecycleState)); !ok && request.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", request.LifecycleState, strings.Join(GetListFindingsLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFindingsReferencesEnum(string(request.References)); !ok && request.References != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for References: %s. Supported values are: %s.", request.References, strings.Join(GetListFindingsReferencesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListFindingsAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListFindingsAccessLevelEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListFindingsResponse wrapper for the ListFindings operation
type ListFindingsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of []FindingSummary instances
	Items []FindingSummary `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListFindingsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListFindingsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListFindingsSeverityEnum Enum with underlying type: string
type ListFindingsSeverityEnum string

// Set of constants representing the allowable values for ListFindingsSeverityEnum
const (
	ListFindingsSeverityHigh     ListFindingsSeverityEnum = "HIGH"
	ListFindingsSeverityMedium   ListFindingsSeverityEnum = "MEDIUM"
	ListFindingsSeverityLow      ListFindingsSeverityEnum = "LOW"
	ListFindingsSeverityEvaluate ListFindingsSeverityEnum = "EVALUATE"
	ListFindingsSeverityAdvisory ListFindingsSeverityEnum = "ADVISORY"
	ListFindingsSeverityPass     ListFindingsSeverityEnum = "PASS"
	ListFindingsSeverityDeferred ListFindingsSeverityEnum = "DEFERRED"
)

var mappingListFindingsSeverityEnum = map[string]ListFindingsSeverityEnum{
	"HIGH":     ListFindingsSeverityHigh,
	"MEDIUM":   ListFindingsSeverityMedium,
	"LOW":      ListFindingsSeverityLow,
	"EVALUATE": ListFindingsSeverityEvaluate,
	"ADVISORY": ListFindingsSeverityAdvisory,
	"PASS":     ListFindingsSeverityPass,
	"DEFERRED": ListFindingsSeverityDeferred,
}

var mappingListFindingsSeverityEnumLowerCase = map[string]ListFindingsSeverityEnum{
	"high":     ListFindingsSeverityHigh,
	"medium":   ListFindingsSeverityMedium,
	"low":      ListFindingsSeverityLow,
	"evaluate": ListFindingsSeverityEvaluate,
	"advisory": ListFindingsSeverityAdvisory,
	"pass":     ListFindingsSeverityPass,
	"deferred": ListFindingsSeverityDeferred,
}

// GetListFindingsSeverityEnumValues Enumerates the set of values for ListFindingsSeverityEnum
func GetListFindingsSeverityEnumValues() []ListFindingsSeverityEnum {
	values := make([]ListFindingsSeverityEnum, 0)
	for _, v := range mappingListFindingsSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingsSeverityEnumStringValues Enumerates the set of values in String for ListFindingsSeverityEnum
func GetListFindingsSeverityEnumStringValues() []string {
	return []string{
		"HIGH",
		"MEDIUM",
		"LOW",
		"EVALUATE",
		"ADVISORY",
		"PASS",
		"DEFERRED",
	}
}

// GetMappingListFindingsSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingsSeverityEnum(val string) (ListFindingsSeverityEnum, bool) {
	enum, ok := mappingListFindingsSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFindingsLifecycleStateEnum Enum with underlying type: string
type ListFindingsLifecycleStateEnum string

// Set of constants representing the allowable values for ListFindingsLifecycleStateEnum
const (
	ListFindingsLifecycleStateActive         ListFindingsLifecycleStateEnum = "ACTIVE"
	ListFindingsLifecycleStateUpdating       ListFindingsLifecycleStateEnum = "UPDATING"
	ListFindingsLifecycleStateNeedsAttention ListFindingsLifecycleStateEnum = "NEEDS_ATTENTION"
	ListFindingsLifecycleStateFailed         ListFindingsLifecycleStateEnum = "FAILED"
)

var mappingListFindingsLifecycleStateEnum = map[string]ListFindingsLifecycleStateEnum{
	"ACTIVE":          ListFindingsLifecycleStateActive,
	"UPDATING":        ListFindingsLifecycleStateUpdating,
	"NEEDS_ATTENTION": ListFindingsLifecycleStateNeedsAttention,
	"FAILED":          ListFindingsLifecycleStateFailed,
}

var mappingListFindingsLifecycleStateEnumLowerCase = map[string]ListFindingsLifecycleStateEnum{
	"active":          ListFindingsLifecycleStateActive,
	"updating":        ListFindingsLifecycleStateUpdating,
	"needs_attention": ListFindingsLifecycleStateNeedsAttention,
	"failed":          ListFindingsLifecycleStateFailed,
}

// GetListFindingsLifecycleStateEnumValues Enumerates the set of values for ListFindingsLifecycleStateEnum
func GetListFindingsLifecycleStateEnumValues() []ListFindingsLifecycleStateEnum {
	values := make([]ListFindingsLifecycleStateEnum, 0)
	for _, v := range mappingListFindingsLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingsLifecycleStateEnumStringValues Enumerates the set of values in String for ListFindingsLifecycleStateEnum
func GetListFindingsLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"UPDATING",
		"NEEDS_ATTENTION",
		"FAILED",
	}
}

// GetMappingListFindingsLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingsLifecycleStateEnum(val string) (ListFindingsLifecycleStateEnum, bool) {
	enum, ok := mappingListFindingsLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFindingsReferencesEnum Enum with underlying type: string
type ListFindingsReferencesEnum string

// Set of constants representing the allowable values for ListFindingsReferencesEnum
const (
	ListFindingsReferencesStig ListFindingsReferencesEnum = "STIG"
	ListFindingsReferencesCis  ListFindingsReferencesEnum = "CIS"
	ListFindingsReferencesGdpr ListFindingsReferencesEnum = "GDPR"
)

var mappingListFindingsReferencesEnum = map[string]ListFindingsReferencesEnum{
	"STIG": ListFindingsReferencesStig,
	"CIS":  ListFindingsReferencesCis,
	"GDPR": ListFindingsReferencesGdpr,
}

var mappingListFindingsReferencesEnumLowerCase = map[string]ListFindingsReferencesEnum{
	"stig": ListFindingsReferencesStig,
	"cis":  ListFindingsReferencesCis,
	"gdpr": ListFindingsReferencesGdpr,
}

// GetListFindingsReferencesEnumValues Enumerates the set of values for ListFindingsReferencesEnum
func GetListFindingsReferencesEnumValues() []ListFindingsReferencesEnum {
	values := make([]ListFindingsReferencesEnum, 0)
	for _, v := range mappingListFindingsReferencesEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingsReferencesEnumStringValues Enumerates the set of values in String for ListFindingsReferencesEnum
func GetListFindingsReferencesEnumStringValues() []string {
	return []string{
		"STIG",
		"CIS",
		"GDPR",
	}
}

// GetMappingListFindingsReferencesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingsReferencesEnum(val string) (ListFindingsReferencesEnum, bool) {
	enum, ok := mappingListFindingsReferencesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListFindingsAccessLevelEnum Enum with underlying type: string
type ListFindingsAccessLevelEnum string

// Set of constants representing the allowable values for ListFindingsAccessLevelEnum
const (
	ListFindingsAccessLevelRestricted ListFindingsAccessLevelEnum = "RESTRICTED"
	ListFindingsAccessLevelAccessible ListFindingsAccessLevelEnum = "ACCESSIBLE"
)

var mappingListFindingsAccessLevelEnum = map[string]ListFindingsAccessLevelEnum{
	"RESTRICTED": ListFindingsAccessLevelRestricted,
	"ACCESSIBLE": ListFindingsAccessLevelAccessible,
}

var mappingListFindingsAccessLevelEnumLowerCase = map[string]ListFindingsAccessLevelEnum{
	"restricted": ListFindingsAccessLevelRestricted,
	"accessible": ListFindingsAccessLevelAccessible,
}

// GetListFindingsAccessLevelEnumValues Enumerates the set of values for ListFindingsAccessLevelEnum
func GetListFindingsAccessLevelEnumValues() []ListFindingsAccessLevelEnum {
	values := make([]ListFindingsAccessLevelEnum, 0)
	for _, v := range mappingListFindingsAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListFindingsAccessLevelEnumStringValues Enumerates the set of values in String for ListFindingsAccessLevelEnum
func GetListFindingsAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListFindingsAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListFindingsAccessLevelEnum(val string) (ListFindingsAccessLevelEnum, bool) {
	enum, ok := mappingListFindingsAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
