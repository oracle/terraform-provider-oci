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

// RemoveQuotaRuleRequest wrapper for the RemoveQuotaRule operation
type RemoveQuotaRuleRequest struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system.
	FileSystemId *string `mandatory:"true" contributesTo:"path" name:"fileSystemId"`

	// An identifier for the owner of this usage and quota rule. Unix-like operating systems use this integer value to
	// identify a user or group to manage access control.
	PrincipalId *int `mandatory:"true" contributesTo:"query" name:"principalId"`

	// The type of the owner of this quota rule and usage.
	PrincipalType RemoveQuotaRulePrincipalTypeEnum `mandatory:"true" contributesTo:"query" name:"principalType" omitEmpty:"true"`

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

func (request RemoveQuotaRuleRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request RemoveQuotaRuleRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request RemoveQuotaRuleRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request RemoveQuotaRuleRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request RemoveQuotaRuleRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveQuotaRulePrincipalTypeEnum(string(request.PrincipalType)); !ok && request.PrincipalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PrincipalType: %s. Supported values are: %s.", request.PrincipalType, strings.Join(GetRemoveQuotaRulePrincipalTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveQuotaRuleResponse wrapper for the RemoveQuotaRule operation
type RemoveQuotaRuleResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// Unique Oracle-assigned identifier for the request. If
	// you need to contact Oracle about a particular request,
	// please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response RemoveQuotaRuleResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response RemoveQuotaRuleResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// RemoveQuotaRulePrincipalTypeEnum Enum with underlying type: string
type RemoveQuotaRulePrincipalTypeEnum string

// Set of constants representing the allowable values for RemoveQuotaRulePrincipalTypeEnum
const (
	RemoveQuotaRulePrincipalTypeFileSystemLevel RemoveQuotaRulePrincipalTypeEnum = "FILE_SYSTEM_LEVEL"
	RemoveQuotaRulePrincipalTypeDefaultGroup    RemoveQuotaRulePrincipalTypeEnum = "DEFAULT_GROUP"
	RemoveQuotaRulePrincipalTypeDefaultUser     RemoveQuotaRulePrincipalTypeEnum = "DEFAULT_USER"
	RemoveQuotaRulePrincipalTypeIndividualGroup RemoveQuotaRulePrincipalTypeEnum = "INDIVIDUAL_GROUP"
	RemoveQuotaRulePrincipalTypeIndividualUser  RemoveQuotaRulePrincipalTypeEnum = "INDIVIDUAL_USER"
)

var mappingRemoveQuotaRulePrincipalTypeEnum = map[string]RemoveQuotaRulePrincipalTypeEnum{
	"FILE_SYSTEM_LEVEL": RemoveQuotaRulePrincipalTypeFileSystemLevel,
	"DEFAULT_GROUP":     RemoveQuotaRulePrincipalTypeDefaultGroup,
	"DEFAULT_USER":      RemoveQuotaRulePrincipalTypeDefaultUser,
	"INDIVIDUAL_GROUP":  RemoveQuotaRulePrincipalTypeIndividualGroup,
	"INDIVIDUAL_USER":   RemoveQuotaRulePrincipalTypeIndividualUser,
}

var mappingRemoveQuotaRulePrincipalTypeEnumLowerCase = map[string]RemoveQuotaRulePrincipalTypeEnum{
	"file_system_level": RemoveQuotaRulePrincipalTypeFileSystemLevel,
	"default_group":     RemoveQuotaRulePrincipalTypeDefaultGroup,
	"default_user":      RemoveQuotaRulePrincipalTypeDefaultUser,
	"individual_group":  RemoveQuotaRulePrincipalTypeIndividualGroup,
	"individual_user":   RemoveQuotaRulePrincipalTypeIndividualUser,
}

// GetRemoveQuotaRulePrincipalTypeEnumValues Enumerates the set of values for RemoveQuotaRulePrincipalTypeEnum
func GetRemoveQuotaRulePrincipalTypeEnumValues() []RemoveQuotaRulePrincipalTypeEnum {
	values := make([]RemoveQuotaRulePrincipalTypeEnum, 0)
	for _, v := range mappingRemoveQuotaRulePrincipalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveQuotaRulePrincipalTypeEnumStringValues Enumerates the set of values in String for RemoveQuotaRulePrincipalTypeEnum
func GetRemoveQuotaRulePrincipalTypeEnumStringValues() []string {
	return []string{
		"FILE_SYSTEM_LEVEL",
		"DEFAULT_GROUP",
		"DEFAULT_USER",
		"INDIVIDUAL_GROUP",
		"INDIVIDUAL_USER",
	}
}

// GetMappingRemoveQuotaRulePrincipalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveQuotaRulePrincipalTypeEnum(val string) (RemoveQuotaRulePrincipalTypeEnum, bool) {
	enum, ok := mappingRemoveQuotaRulePrincipalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
