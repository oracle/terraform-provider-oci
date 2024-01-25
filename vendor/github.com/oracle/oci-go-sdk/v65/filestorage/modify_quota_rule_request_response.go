// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"net/http"
	"strings"
)

// ModifyQuotaRuleRequest wrapper for the ModifyQuotaRule operation
type ModifyQuotaRuleRequest struct {

	// Details for editting a new quota rule.
	ModifyQuotaRuleDetails `contributesTo:"body"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`

	// An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"true" contributesTo:"query" name:"principalId"`

	// The type of the owner of this quota rule and usage.
	PrincipalType ModifyQuotaRulePrincipalTypeEnum `mandatory:"true" contributesTo:"query" name:"principalType" omitEmpty:"true"`

	// The flag is an identifier to tell whether the quota rule will be enforced.
	// If `isHardQuota` is false, the quota rule will be enforced so the usage cannot exceed the hard quota limit.
	// If `isHardQuota` is true, usage can exceed the soft quota limit. An alarm or notification will be sent to
	// the customer, if the specific usage exceeds.
	IsHardQuota *bool `mandatory:"true" contributesTo:"query" name:"isHardQuota"`

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

func (request ModifyQuotaRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ModifyQuotaRuleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ModifyQuotaRuleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ModifyQuotaRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ModifyQuotaRuleRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModifyQuotaRulePrincipalTypeEnum(string(request.PrincipalType)); !ok && request.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", request.PrincipalType, strings.Join(GetModifyQuotaRulePrincipalTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModifyQuotaRuleResponse wrapper for the ModifyQuotaRule operation
type ModifyQuotaRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The QuotaRule instance
	QuotaRule `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ModifyQuotaRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ModifyQuotaRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ModifyQuotaRulePrincipalTypeEnum Enum with underlying type: string
type ModifyQuotaRulePrincipalTypeEnum string

// Set of constants representing the allowable values for ModifyQuotaRulePrincipalTypeEnum
const (
	ModifyQuotaRulePrincipalTypeFileSystemLevel ModifyQuotaRulePrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	ModifyQuotaRulePrincipalTypeDefaultGroup    ModifyQuotaRulePrincipalTypeEnum = "DEFAULT_GROUP"
	ModifyQuotaRulePrincipalTypeDefaultUser     ModifyQuotaRulePrincipalTypeEnum = "DEFAULT_USER"
	ModifyQuotaRulePrincipalTypeIndividualGroup ModifyQuotaRulePrincipalTypeEnum = "INDIVIDUAL_GROUP"
	ModifyQuotaRulePrincipalTypeIndividualUser  ModifyQuotaRulePrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingModifyQuotaRulePrincipalTypeEnum = map[string]ModifyQuotaRulePrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": ModifyQuotaRulePrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     ModifyQuotaRulePrincipalTypeDefaultGroup,
	"DEFAULT_USER":      ModifyQuotaRulePrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  ModifyQuotaRulePrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   ModifyQuotaRulePrincipalTypeIndividualUser,
}

var mappingModifyQuotaRulePrincipalTypeEnumLowerCase = map[string]ModifyQuotaRulePrincipalTypeEnum{
	"file_system_level": ModifyQuotaRulePrincipalTypeFileSystemLevel,
	"default_group":     ModifyQuotaRulePrincipalTypeDefaultGroup,
	"default_user":      ModifyQuotaRulePrincipalTypeDefaultUser,
	"individual_group":  ModifyQuotaRulePrincipalTypeIndividualGroup,
	"individual_user":   ModifyQuotaRulePrincipalTypeIndividualUser,
}

// GetModifyQuotaRulePrincipalTypeEnumValues Enumerates the set of values for ModifyQuotaRulePrincipalTypeEnum
func GetModifyQuotaRulePrincipalTypeEnumValues() []ModifyQuotaRulePrincipalTypeEnum {
	values := make([]ModifyQuotaRulePrincipalTypeEnum, 0)
	for _, v := range mappingModifyQuotaRulePrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetModifyQuotaRulePrincipalTypeEnumStringValues Enumerates the set of values in String for ModifyQuotaRulePrincipalTypeEnum
func GetModifyQuotaRulePrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingModifyQuotaRulePrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifyQuotaRulePrincipalTypeEnum(val string) (ModifyQuotaRulePrincipalTypeEnum, bool) {
	enum, ok := mappingModifyQuotaRulePrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
