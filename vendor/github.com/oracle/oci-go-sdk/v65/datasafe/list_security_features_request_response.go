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

// ListSecurityFeaturesRequest wrapper for the ListSecurityFeatures operation
//
// # See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/datasafe/ListSecurityFeatures.go.html to see an example of how to use ListSecurityFeaturesRequest.
type ListSecurityFeaturesRequest struct {

	// A filter to return only resources that match the specified compartment OCID.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// Default is false.
	// When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting.
	CompartmentIdInSubtree *bool `mandatory:"false" contributesTo:"query" name:"compartmentIdInSubtree"`

	// Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED.
	// Setting this to ACCESSIBLE returns only those compartments for which the
	// user has INSPECT permissions directly or indirectly (permissions can be on a
	// resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed.
	AccessLevel ListSecurityFeaturesAccessLevelEnum `mandatory:"false" contributesTo:"query" name:"accessLevel" omitEmpty:"true"`

	// A filter to return only items related to a specific target OCID.
	TargetId *string `mandatory:"false" contributesTo:"query" name:"targetId"`

	// A filter to return only the targets with the DB security feature - Unified Audit enabled/disabled.
	TargetsWithUnifiedAudit ListSecurityFeaturesTargetsWithUnifiedAuditEnum `mandatory:"false" contributesTo:"query" name:"targetsWithUnifiedAudit" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Fine Grained Audit enabled/disabled.
	TargetsWithFineGrainedAudit ListSecurityFeaturesTargetsWithFineGrainedAuditEnum `mandatory:"false" contributesTo:"query" name:"targetsWithFineGrainedAudit" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Traditional Audit enabled/disabled.
	TargetsWithTraditionalAudit ListSecurityFeaturesTargetsWithTraditionalAuditEnum `mandatory:"false" contributesTo:"query" name:"targetsWithTraditionalAudit" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Database Vault enabled/disabled.
	TargetsWithDatabaseVault ListSecurityFeaturesTargetsWithDatabaseVaultEnum `mandatory:"false" contributesTo:"query" name:"targetsWithDatabaseVault" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Privilege Analysis enabled/disabled.
	TargetsWithPrivilegeAnalysis ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum `mandatory:"false" contributesTo:"query" name:"targetsWithPrivilegeAnalysis" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Tablespace Encryption enabled/disabled.
	TargetsWithTablespaceEncryption ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum `mandatory:"false" contributesTo:"query" name:"targetsWithTablespaceEncryption" omitEmpty:"true"`

	// A filter to return only the targets that enable the DB security feature - Column Encryption enabled/disabled.
	TargetsWithColumnEncryption ListSecurityFeaturesTargetsWithColumnEncryptionEnum `mandatory:"false" contributesTo:"query" name:"targetsWithColumnEncryption" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Network Encryption enabled/disabled.
	TargetsWithNetworkEncryption ListSecurityFeaturesTargetsWithNetworkEncryptionEnum `mandatory:"false" contributesTo:"query" name:"targetsWithNetworkEncryption" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Password Authentication enabled/disabled.
	TargetsWithPasswordAuthentication ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum `mandatory:"false" contributesTo:"query" name:"targetsWithPasswordAuthentication" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - Global Authentication enabled/disabled.
	TargetsWithGlobalAuthentication ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum `mandatory:"false" contributesTo:"query" name:"targetsWithGlobalAuthentication" omitEmpty:"true"`

	// A filter to return only the targets with the DB security feature - External Authentication enabled/disabled.
	TargetsWithExternalAuthentication ListSecurityFeaturesTargetsWithExternalAuthenticationEnum `mandatory:"false" contributesTo:"query" name:"targetsWithExternalAuthentication" omitEmpty:"true"`

	// Unique identifier for the request.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// For list pagination. The maximum number of items to return per page in a paginated "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// For list pagination. The page token representing the page at which to start retrieving results. It is usually retrieved from a previous "List" call. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/en-us/iaas/Content/API/Concepts/usingapi.htm#nine).
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListSecurityFeaturesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListSecurityFeaturesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListSecurityFeaturesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListSecurityFeaturesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListSecurityFeaturesRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListSecurityFeaturesAccessLevelEnum(string(request.AccessLevel)); !ok && request.AccessLevel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AccessLevel: %s. Supported values are: %s.", request.AccessLevel, strings.Join(GetListSecurityFeaturesAccessLevelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithUnifiedAuditEnum(string(request.TargetsWithUnifiedAudit)); !ok && request.TargetsWithUnifiedAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithUnifiedAudit: %s. Supported values are: %s.", request.TargetsWithUnifiedAudit, strings.Join(GetListSecurityFeaturesTargetsWithUnifiedAuditEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithFineGrainedAuditEnum(string(request.TargetsWithFineGrainedAudit)); !ok && request.TargetsWithFineGrainedAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithFineGrainedAudit: %s. Supported values are: %s.", request.TargetsWithFineGrainedAudit, strings.Join(GetListSecurityFeaturesTargetsWithFineGrainedAuditEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithTraditionalAuditEnum(string(request.TargetsWithTraditionalAudit)); !ok && request.TargetsWithTraditionalAudit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithTraditionalAudit: %s. Supported values are: %s.", request.TargetsWithTraditionalAudit, strings.Join(GetListSecurityFeaturesTargetsWithTraditionalAuditEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithDatabaseVaultEnum(string(request.TargetsWithDatabaseVault)); !ok && request.TargetsWithDatabaseVault != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithDatabaseVault: %s. Supported values are: %s.", request.TargetsWithDatabaseVault, strings.Join(GetListSecurityFeaturesTargetsWithDatabaseVaultEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum(string(request.TargetsWithPrivilegeAnalysis)); !ok && request.TargetsWithPrivilegeAnalysis != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithPrivilegeAnalysis: %s. Supported values are: %s.", request.TargetsWithPrivilegeAnalysis, strings.Join(GetListSecurityFeaturesTargetsWithPrivilegeAnalysisEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithTablespaceEncryptionEnum(string(request.TargetsWithTablespaceEncryption)); !ok && request.TargetsWithTablespaceEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithTablespaceEncryption: %s. Supported values are: %s.", request.TargetsWithTablespaceEncryption, strings.Join(GetListSecurityFeaturesTargetsWithTablespaceEncryptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithColumnEncryptionEnum(string(request.TargetsWithColumnEncryption)); !ok && request.TargetsWithColumnEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithColumnEncryption: %s. Supported values are: %s.", request.TargetsWithColumnEncryption, strings.Join(GetListSecurityFeaturesTargetsWithColumnEncryptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithNetworkEncryptionEnum(string(request.TargetsWithNetworkEncryption)); !ok && request.TargetsWithNetworkEncryption != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithNetworkEncryption: %s. Supported values are: %s.", request.TargetsWithNetworkEncryption, strings.Join(GetListSecurityFeaturesTargetsWithNetworkEncryptionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithPasswordAuthenticationEnum(string(request.TargetsWithPasswordAuthentication)); !ok && request.TargetsWithPasswordAuthentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithPasswordAuthentication: %s. Supported values are: %s.", request.TargetsWithPasswordAuthentication, strings.Join(GetListSecurityFeaturesTargetsWithPasswordAuthenticationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithGlobalAuthenticationEnum(string(request.TargetsWithGlobalAuthentication)); !ok && request.TargetsWithGlobalAuthentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithGlobalAuthentication: %s. Supported values are: %s.", request.TargetsWithGlobalAuthentication, strings.Join(GetListSecurityFeaturesTargetsWithGlobalAuthenticationEnumStringValues(), ",")))
	}
	if _, ok := GetMappingListSecurityFeaturesTargetsWithExternalAuthenticationEnum(string(request.TargetsWithExternalAuthentication)); !ok && request.TargetsWithExternalAuthentication != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetsWithExternalAuthentication: %s. Supported values are: %s.", request.TargetsWithExternalAuthentication, strings.Join(GetListSecurityFeaturesTargetsWithExternalAuthenticationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListSecurityFeaturesResponse wrapper for the ListSecurityFeatures operation
type ListSecurityFeaturesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of SecurityFeatureCollection instances
	SecurityFeatureCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For list pagination. When this header appears in the response, additional pages of results remain. Include opc-next-page value as the page parameter for the subsequent GET request to get the next batch of items. For details about how pagination works, see List Pagination (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm#nine).
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the previous batch of items.
	OpcPrevPage *string `presentIn:"header" name:"opc-prev-page"`
}

func (response ListSecurityFeaturesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListSecurityFeaturesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListSecurityFeaturesAccessLevelEnum Enum with underlying type: string
type ListSecurityFeaturesAccessLevelEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesAccessLevelEnum
const (
	ListSecurityFeaturesAccessLevelRestricted ListSecurityFeaturesAccessLevelEnum = "RESTRICTED"
	ListSecurityFeaturesAccessLevelAccessible ListSecurityFeaturesAccessLevelEnum = "ACCESSIBLE"
)

var mappingListSecurityFeaturesAccessLevelEnum = map[string]ListSecurityFeaturesAccessLevelEnum{
	"RESTRICTED": ListSecurityFeaturesAccessLevelRestricted,
	"ACCESSIBLE": ListSecurityFeaturesAccessLevelAccessible,
}

var mappingListSecurityFeaturesAccessLevelEnumLowerCase = map[string]ListSecurityFeaturesAccessLevelEnum{
	"restricted": ListSecurityFeaturesAccessLevelRestricted,
	"accessible": ListSecurityFeaturesAccessLevelAccessible,
}

// GetListSecurityFeaturesAccessLevelEnumValues Enumerates the set of values for ListSecurityFeaturesAccessLevelEnum
func GetListSecurityFeaturesAccessLevelEnumValues() []ListSecurityFeaturesAccessLevelEnum {
	values := make([]ListSecurityFeaturesAccessLevelEnum, 0)
	for _, v := range mappingListSecurityFeaturesAccessLevelEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesAccessLevelEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesAccessLevelEnum
func GetListSecurityFeaturesAccessLevelEnumStringValues() []string {
	return []string{
		"RESTRICTED",
		"ACCESSIBLE",
	}
}

// GetMappingListSecurityFeaturesAccessLevelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesAccessLevelEnum(val string) (ListSecurityFeaturesAccessLevelEnum, bool) {
	enum, ok := mappingListSecurityFeaturesAccessLevelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithUnifiedAuditEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithUnifiedAuditEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithUnifiedAuditEnum
const (
	ListSecurityFeaturesTargetsWithUnifiedAuditEnabled  ListSecurityFeaturesTargetsWithUnifiedAuditEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithUnifiedAuditDisabled ListSecurityFeaturesTargetsWithUnifiedAuditEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithUnifiedAuditNone     ListSecurityFeaturesTargetsWithUnifiedAuditEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithUnifiedAuditEnum = map[string]ListSecurityFeaturesTargetsWithUnifiedAuditEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithUnifiedAuditEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithUnifiedAuditDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithUnifiedAuditNone,
}

var mappingListSecurityFeaturesTargetsWithUnifiedAuditEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithUnifiedAuditEnum{
	"enabled":  ListSecurityFeaturesTargetsWithUnifiedAuditEnabled,
	"disabled": ListSecurityFeaturesTargetsWithUnifiedAuditDisabled,
	"none":     ListSecurityFeaturesTargetsWithUnifiedAuditNone,
}

// GetListSecurityFeaturesTargetsWithUnifiedAuditEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithUnifiedAuditEnum
func GetListSecurityFeaturesTargetsWithUnifiedAuditEnumValues() []ListSecurityFeaturesTargetsWithUnifiedAuditEnum {
	values := make([]ListSecurityFeaturesTargetsWithUnifiedAuditEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithUnifiedAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithUnifiedAuditEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithUnifiedAuditEnum
func GetListSecurityFeaturesTargetsWithUnifiedAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithUnifiedAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithUnifiedAuditEnum(val string) (ListSecurityFeaturesTargetsWithUnifiedAuditEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithUnifiedAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithFineGrainedAuditEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithFineGrainedAuditEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithFineGrainedAuditEnum
const (
	ListSecurityFeaturesTargetsWithFineGrainedAuditEnabled  ListSecurityFeaturesTargetsWithFineGrainedAuditEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithFineGrainedAuditDisabled ListSecurityFeaturesTargetsWithFineGrainedAuditEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithFineGrainedAuditNone     ListSecurityFeaturesTargetsWithFineGrainedAuditEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithFineGrainedAuditEnum = map[string]ListSecurityFeaturesTargetsWithFineGrainedAuditEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithFineGrainedAuditEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithFineGrainedAuditDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithFineGrainedAuditNone,
}

var mappingListSecurityFeaturesTargetsWithFineGrainedAuditEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithFineGrainedAuditEnum{
	"enabled":  ListSecurityFeaturesTargetsWithFineGrainedAuditEnabled,
	"disabled": ListSecurityFeaturesTargetsWithFineGrainedAuditDisabled,
	"none":     ListSecurityFeaturesTargetsWithFineGrainedAuditNone,
}

// GetListSecurityFeaturesTargetsWithFineGrainedAuditEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithFineGrainedAuditEnum
func GetListSecurityFeaturesTargetsWithFineGrainedAuditEnumValues() []ListSecurityFeaturesTargetsWithFineGrainedAuditEnum {
	values := make([]ListSecurityFeaturesTargetsWithFineGrainedAuditEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithFineGrainedAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithFineGrainedAuditEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithFineGrainedAuditEnum
func GetListSecurityFeaturesTargetsWithFineGrainedAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithFineGrainedAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithFineGrainedAuditEnum(val string) (ListSecurityFeaturesTargetsWithFineGrainedAuditEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithFineGrainedAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithTraditionalAuditEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithTraditionalAuditEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithTraditionalAuditEnum
const (
	ListSecurityFeaturesTargetsWithTraditionalAuditEnabled  ListSecurityFeaturesTargetsWithTraditionalAuditEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithTraditionalAuditDisabled ListSecurityFeaturesTargetsWithTraditionalAuditEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithTraditionalAuditNone     ListSecurityFeaturesTargetsWithTraditionalAuditEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithTraditionalAuditEnum = map[string]ListSecurityFeaturesTargetsWithTraditionalAuditEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithTraditionalAuditEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithTraditionalAuditDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithTraditionalAuditNone,
}

var mappingListSecurityFeaturesTargetsWithTraditionalAuditEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithTraditionalAuditEnum{
	"enabled":  ListSecurityFeaturesTargetsWithTraditionalAuditEnabled,
	"disabled": ListSecurityFeaturesTargetsWithTraditionalAuditDisabled,
	"none":     ListSecurityFeaturesTargetsWithTraditionalAuditNone,
}

// GetListSecurityFeaturesTargetsWithTraditionalAuditEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithTraditionalAuditEnum
func GetListSecurityFeaturesTargetsWithTraditionalAuditEnumValues() []ListSecurityFeaturesTargetsWithTraditionalAuditEnum {
	values := make([]ListSecurityFeaturesTargetsWithTraditionalAuditEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithTraditionalAuditEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithTraditionalAuditEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithTraditionalAuditEnum
func GetListSecurityFeaturesTargetsWithTraditionalAuditEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithTraditionalAuditEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithTraditionalAuditEnum(val string) (ListSecurityFeaturesTargetsWithTraditionalAuditEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithTraditionalAuditEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithDatabaseVaultEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithDatabaseVaultEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithDatabaseVaultEnum
const (
	ListSecurityFeaturesTargetsWithDatabaseVaultEnabled  ListSecurityFeaturesTargetsWithDatabaseVaultEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithDatabaseVaultDisabled ListSecurityFeaturesTargetsWithDatabaseVaultEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithDatabaseVaultNone     ListSecurityFeaturesTargetsWithDatabaseVaultEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithDatabaseVaultEnum = map[string]ListSecurityFeaturesTargetsWithDatabaseVaultEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithDatabaseVaultEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithDatabaseVaultDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithDatabaseVaultNone,
}

var mappingListSecurityFeaturesTargetsWithDatabaseVaultEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithDatabaseVaultEnum{
	"enabled":  ListSecurityFeaturesTargetsWithDatabaseVaultEnabled,
	"disabled": ListSecurityFeaturesTargetsWithDatabaseVaultDisabled,
	"none":     ListSecurityFeaturesTargetsWithDatabaseVaultNone,
}

// GetListSecurityFeaturesTargetsWithDatabaseVaultEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithDatabaseVaultEnum
func GetListSecurityFeaturesTargetsWithDatabaseVaultEnumValues() []ListSecurityFeaturesTargetsWithDatabaseVaultEnum {
	values := make([]ListSecurityFeaturesTargetsWithDatabaseVaultEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithDatabaseVaultEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithDatabaseVaultEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithDatabaseVaultEnum
func GetListSecurityFeaturesTargetsWithDatabaseVaultEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithDatabaseVaultEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithDatabaseVaultEnum(val string) (ListSecurityFeaturesTargetsWithDatabaseVaultEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithDatabaseVaultEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum
const (
	ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnabled  ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithPrivilegeAnalysisDisabled ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithPrivilegeAnalysisNone     ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum = map[string]ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithPrivilegeAnalysisDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithPrivilegeAnalysisNone,
}

var mappingListSecurityFeaturesTargetsWithPrivilegeAnalysisEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum{
	"enabled":  ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnabled,
	"disabled": ListSecurityFeaturesTargetsWithPrivilegeAnalysisDisabled,
	"none":     ListSecurityFeaturesTargetsWithPrivilegeAnalysisNone,
}

// GetListSecurityFeaturesTargetsWithPrivilegeAnalysisEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum
func GetListSecurityFeaturesTargetsWithPrivilegeAnalysisEnumValues() []ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum {
	values := make([]ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithPrivilegeAnalysisEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum
func GetListSecurityFeaturesTargetsWithPrivilegeAnalysisEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum(val string) (ListSecurityFeaturesTargetsWithPrivilegeAnalysisEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithPrivilegeAnalysisEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum
const (
	ListSecurityFeaturesTargetsWithTablespaceEncryptionEnabled  ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithTablespaceEncryptionDisabled ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithTablespaceEncryptionNone     ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithTablespaceEncryptionEnum = map[string]ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithTablespaceEncryptionEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithTablespaceEncryptionDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithTablespaceEncryptionNone,
}

var mappingListSecurityFeaturesTargetsWithTablespaceEncryptionEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum{
	"enabled":  ListSecurityFeaturesTargetsWithTablespaceEncryptionEnabled,
	"disabled": ListSecurityFeaturesTargetsWithTablespaceEncryptionDisabled,
	"none":     ListSecurityFeaturesTargetsWithTablespaceEncryptionNone,
}

// GetListSecurityFeaturesTargetsWithTablespaceEncryptionEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum
func GetListSecurityFeaturesTargetsWithTablespaceEncryptionEnumValues() []ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum {
	values := make([]ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithTablespaceEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithTablespaceEncryptionEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum
func GetListSecurityFeaturesTargetsWithTablespaceEncryptionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithTablespaceEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithTablespaceEncryptionEnum(val string) (ListSecurityFeaturesTargetsWithTablespaceEncryptionEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithTablespaceEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithColumnEncryptionEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithColumnEncryptionEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithColumnEncryptionEnum
const (
	ListSecurityFeaturesTargetsWithColumnEncryptionEnabled  ListSecurityFeaturesTargetsWithColumnEncryptionEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithColumnEncryptionDisabled ListSecurityFeaturesTargetsWithColumnEncryptionEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithColumnEncryptionNone     ListSecurityFeaturesTargetsWithColumnEncryptionEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithColumnEncryptionEnum = map[string]ListSecurityFeaturesTargetsWithColumnEncryptionEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithColumnEncryptionEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithColumnEncryptionDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithColumnEncryptionNone,
}

var mappingListSecurityFeaturesTargetsWithColumnEncryptionEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithColumnEncryptionEnum{
	"enabled":  ListSecurityFeaturesTargetsWithColumnEncryptionEnabled,
	"disabled": ListSecurityFeaturesTargetsWithColumnEncryptionDisabled,
	"none":     ListSecurityFeaturesTargetsWithColumnEncryptionNone,
}

// GetListSecurityFeaturesTargetsWithColumnEncryptionEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithColumnEncryptionEnum
func GetListSecurityFeaturesTargetsWithColumnEncryptionEnumValues() []ListSecurityFeaturesTargetsWithColumnEncryptionEnum {
	values := make([]ListSecurityFeaturesTargetsWithColumnEncryptionEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithColumnEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithColumnEncryptionEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithColumnEncryptionEnum
func GetListSecurityFeaturesTargetsWithColumnEncryptionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithColumnEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithColumnEncryptionEnum(val string) (ListSecurityFeaturesTargetsWithColumnEncryptionEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithColumnEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithNetworkEncryptionEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithNetworkEncryptionEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithNetworkEncryptionEnum
const (
	ListSecurityFeaturesTargetsWithNetworkEncryptionEnabled  ListSecurityFeaturesTargetsWithNetworkEncryptionEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithNetworkEncryptionDisabled ListSecurityFeaturesTargetsWithNetworkEncryptionEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithNetworkEncryptionNone     ListSecurityFeaturesTargetsWithNetworkEncryptionEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithNetworkEncryptionEnum = map[string]ListSecurityFeaturesTargetsWithNetworkEncryptionEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithNetworkEncryptionEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithNetworkEncryptionDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithNetworkEncryptionNone,
}

var mappingListSecurityFeaturesTargetsWithNetworkEncryptionEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithNetworkEncryptionEnum{
	"enabled":  ListSecurityFeaturesTargetsWithNetworkEncryptionEnabled,
	"disabled": ListSecurityFeaturesTargetsWithNetworkEncryptionDisabled,
	"none":     ListSecurityFeaturesTargetsWithNetworkEncryptionNone,
}

// GetListSecurityFeaturesTargetsWithNetworkEncryptionEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithNetworkEncryptionEnum
func GetListSecurityFeaturesTargetsWithNetworkEncryptionEnumValues() []ListSecurityFeaturesTargetsWithNetworkEncryptionEnum {
	values := make([]ListSecurityFeaturesTargetsWithNetworkEncryptionEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithNetworkEncryptionEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithNetworkEncryptionEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithNetworkEncryptionEnum
func GetListSecurityFeaturesTargetsWithNetworkEncryptionEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithNetworkEncryptionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithNetworkEncryptionEnum(val string) (ListSecurityFeaturesTargetsWithNetworkEncryptionEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithNetworkEncryptionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum
const (
	ListSecurityFeaturesTargetsWithPasswordAuthenticationEnabled  ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithPasswordAuthenticationDisabled ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithPasswordAuthenticationNone     ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithPasswordAuthenticationEnum = map[string]ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithPasswordAuthenticationEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithPasswordAuthenticationDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithPasswordAuthenticationNone,
}

var mappingListSecurityFeaturesTargetsWithPasswordAuthenticationEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum{
	"enabled":  ListSecurityFeaturesTargetsWithPasswordAuthenticationEnabled,
	"disabled": ListSecurityFeaturesTargetsWithPasswordAuthenticationDisabled,
	"none":     ListSecurityFeaturesTargetsWithPasswordAuthenticationNone,
}

// GetListSecurityFeaturesTargetsWithPasswordAuthenticationEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum
func GetListSecurityFeaturesTargetsWithPasswordAuthenticationEnumValues() []ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum {
	values := make([]ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithPasswordAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithPasswordAuthenticationEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum
func GetListSecurityFeaturesTargetsWithPasswordAuthenticationEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithPasswordAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithPasswordAuthenticationEnum(val string) (ListSecurityFeaturesTargetsWithPasswordAuthenticationEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithPasswordAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum
const (
	ListSecurityFeaturesTargetsWithGlobalAuthenticationEnabled  ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithGlobalAuthenticationDisabled ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithGlobalAuthenticationNone     ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithGlobalAuthenticationEnum = map[string]ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithGlobalAuthenticationEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithGlobalAuthenticationDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithGlobalAuthenticationNone,
}

var mappingListSecurityFeaturesTargetsWithGlobalAuthenticationEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum{
	"enabled":  ListSecurityFeaturesTargetsWithGlobalAuthenticationEnabled,
	"disabled": ListSecurityFeaturesTargetsWithGlobalAuthenticationDisabled,
	"none":     ListSecurityFeaturesTargetsWithGlobalAuthenticationNone,
}

// GetListSecurityFeaturesTargetsWithGlobalAuthenticationEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum
func GetListSecurityFeaturesTargetsWithGlobalAuthenticationEnumValues() []ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum {
	values := make([]ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithGlobalAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithGlobalAuthenticationEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum
func GetListSecurityFeaturesTargetsWithGlobalAuthenticationEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithGlobalAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithGlobalAuthenticationEnum(val string) (ListSecurityFeaturesTargetsWithGlobalAuthenticationEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithGlobalAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ListSecurityFeaturesTargetsWithExternalAuthenticationEnum Enum with underlying type: string
type ListSecurityFeaturesTargetsWithExternalAuthenticationEnum string

// Set of constants representing the allowable values for ListSecurityFeaturesTargetsWithExternalAuthenticationEnum
const (
	ListSecurityFeaturesTargetsWithExternalAuthenticationEnabled  ListSecurityFeaturesTargetsWithExternalAuthenticationEnum = "ENABLED"
	ListSecurityFeaturesTargetsWithExternalAuthenticationDisabled ListSecurityFeaturesTargetsWithExternalAuthenticationEnum = "DISABLED"
	ListSecurityFeaturesTargetsWithExternalAuthenticationNone     ListSecurityFeaturesTargetsWithExternalAuthenticationEnum = "NONE"
)

var mappingListSecurityFeaturesTargetsWithExternalAuthenticationEnum = map[string]ListSecurityFeaturesTargetsWithExternalAuthenticationEnum{
	"ENABLED":  ListSecurityFeaturesTargetsWithExternalAuthenticationEnabled,
	"DISABLED": ListSecurityFeaturesTargetsWithExternalAuthenticationDisabled,
	"NONE":     ListSecurityFeaturesTargetsWithExternalAuthenticationNone,
}

var mappingListSecurityFeaturesTargetsWithExternalAuthenticationEnumLowerCase = map[string]ListSecurityFeaturesTargetsWithExternalAuthenticationEnum{
	"enabled":  ListSecurityFeaturesTargetsWithExternalAuthenticationEnabled,
	"disabled": ListSecurityFeaturesTargetsWithExternalAuthenticationDisabled,
	"none":     ListSecurityFeaturesTargetsWithExternalAuthenticationNone,
}

// GetListSecurityFeaturesTargetsWithExternalAuthenticationEnumValues Enumerates the set of values for ListSecurityFeaturesTargetsWithExternalAuthenticationEnum
func GetListSecurityFeaturesTargetsWithExternalAuthenticationEnumValues() []ListSecurityFeaturesTargetsWithExternalAuthenticationEnum {
	values := make([]ListSecurityFeaturesTargetsWithExternalAuthenticationEnum, 0)
	for _, v := range mappingListSecurityFeaturesTargetsWithExternalAuthenticationEnum {
		values = append(values, v)
	}
	return values
}

// GetListSecurityFeaturesTargetsWithExternalAuthenticationEnumStringValues Enumerates the set of values in String for ListSecurityFeaturesTargetsWithExternalAuthenticationEnum
func GetListSecurityFeaturesTargetsWithExternalAuthenticationEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NONE",
	}
}

// GetMappingListSecurityFeaturesTargetsWithExternalAuthenticationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListSecurityFeaturesTargetsWithExternalAuthenticationEnum(val string) (ListSecurityFeaturesTargetsWithExternalAuthenticationEnum, bool) {
	enum, ok := mappingListSecurityFeaturesTargetsWithExternalAuthenticationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
