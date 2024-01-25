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

// GetQuotaRuleRequest wrapper for the GetQuotaRule operation
type GetQuotaRuleRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`

	// An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"true" contributesTo:"query" name:"principalId"`

	// The type of the owner of this quota rule and usage.
	PrincipalType GetQuotaRulePrincipalTypeEnum `mandatory:"true" contributesTo:"query" name:"principalType" omitEmpty:"true"`

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

func (request GetQuotaRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request GetQuotaRuleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request GetQuotaRuleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request GetQuotaRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request GetQuotaRuleRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingGetQuotaRulePrincipalTypeEnum(string(request.PrincipalType)); !ok && request.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", request.PrincipalType, strings.Join(GetGetQuotaRulePrincipalTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// GetQuotaRuleResponse wrapper for the GetQuotaRule operation
type GetQuotaRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The QuotaRule instance
	QuotaRule `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetQuotaRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response GetQuotaRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// GetQuotaRulePrincipalTypeEnum Enum with underlying type: string
type GetQuotaRulePrincipalTypeEnum string

// Set of constants representing the allowable values for GetQuotaRulePrincipalTypeEnum
const (
	GetQuotaRulePrincipalTypeFileSystemLevel GetQuotaRulePrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	GetQuotaRulePrincipalTypeDefaultGroup    GetQuotaRulePrincipalTypeEnum = "DEFAULT_GROUP"
	GetQuotaRulePrincipalTypeDefaultUser     GetQuotaRulePrincipalTypeEnum = "DEFAULT_USER"
	GetQuotaRulePrincipalTypeIndividualGroup GetQuotaRulePrincipalTypeEnum = "INDIVIDUAL_GROUP"
	GetQuotaRulePrincipalTypeIndividualUser  GetQuotaRulePrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingGetQuotaRulePrincipalTypeEnum = map[string]GetQuotaRulePrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": GetQuotaRulePrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     GetQuotaRulePrincipalTypeDefaultGroup,
	"DEFAULT_USER":      GetQuotaRulePrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  GetQuotaRulePrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   GetQuotaRulePrincipalTypeIndividualUser,
}

var mappingGetQuotaRulePrincipalTypeEnumLowerCase = map[string]GetQuotaRulePrincipalTypeEnum{
	"file_system_level": GetQuotaRulePrincipalTypeFileSystemLevel,
	"default_group":     GetQuotaRulePrincipalTypeDefaultGroup,
	"default_user":      GetQuotaRulePrincipalTypeDefaultUser,
	"individual_group":  GetQuotaRulePrincipalTypeIndividualGroup,
	"individual_user":   GetQuotaRulePrincipalTypeIndividualUser,
}

// GetGetQuotaRulePrincipalTypeEnumValues Enumerates the set of values for GetQuotaRulePrincipalTypeEnum
func GetGetQuotaRulePrincipalTypeEnumValues() []GetQuotaRulePrincipalTypeEnum {
	values := make([]GetQuotaRulePrincipalTypeEnum, 0)
	for _, v := range mappingGetQuotaRulePrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetGetQuotaRulePrincipalTypeEnumStringValues Enumerates the set of values in String for GetQuotaRulePrincipalTypeEnum
func GetGetQuotaRulePrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingGetQuotaRulePrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingGetQuotaRulePrincipalTypeEnum(val string) (GetQuotaRulePrincipalTypeEnum, bool) {
	enum, ok := mappingGetQuotaRulePrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
