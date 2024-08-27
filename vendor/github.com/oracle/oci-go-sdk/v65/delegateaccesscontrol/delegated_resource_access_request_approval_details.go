// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Delegate Access Control API
//
// Oracle Delegate Access Control allows ExaCC and ExaCS customers to delegate management of their Exadata resources operators outside their tenancies.
// With Delegate Access Control, Support Providers can deliver managed services using comprehensive and robust tooling built on the OCI platform.
// Customers maintain control over who has access to the delegated resources in their tenancy and what actions can be taken.
// Enterprises managing resources across multiple tenants can use Delegate Access Control to streamline management tasks.
// Using logging service, customers can view a near real-time audit report of all actions performed by a Service Provider operator.
//

package delegateaccesscontrol

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DelegatedResourceAccessRequestApprovalDetails Approval info for initial access or extension of a Delegated Resource Access Request
type DelegatedResourceAccessRequestApprovalDetails struct {

	// Comment specified by the approver of the request.
	ApproverComment *string `mandatory:"true" json:"approverComment"`

	// Indicated whether the request is approved or rejected.
	ApprovalAction DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum `mandatory:"false" json:"approvalAction,omitempty"`

	// Access start time that is actually approved by the customer in RFC 3339 (https://tools.ietf.org/html/rfc3339)timestamp format, e.g. '2020-05-22T21:10:29.600Z'.
	TimeApprovedForAccess *common.SDKTime `mandatory:"false" json:"timeApprovedForAccess"`

	// approval type, initial or extension
	ApprovalType DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum `mandatory:"false" json:"approvalType,omitempty"`

	// User ID of the approver.
	ApproverId *string `mandatory:"false" json:"approverId"`

	// Additional message specified by the approver of the request.
	ApproverAdditionalMessage *string `mandatory:"false" json:"approverAdditionalMessage"`
}

func (m DelegatedResourceAccessRequestApprovalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DelegatedResourceAccessRequestApprovalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum(string(m.ApprovalAction)); !ok && m.ApprovalAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApprovalAction: %s. Supported values are: %s.", m.ApprovalAction, strings.Join(GetDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum(string(m.ApprovalType)); !ok && m.ApprovalType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ApprovalType: %s. Supported values are: %s.", m.ApprovalType, strings.Join(GetDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum Enum with underlying type: string
type DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum string

// Set of constants representing the allowable values for DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum
const (
	DelegatedResourceAccessRequestApprovalDetailsApprovalActionApprove DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum = "APPROVE"
	DelegatedResourceAccessRequestApprovalDetailsApprovalActionReject  DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum = "REJECT"
)

var mappingDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum = map[string]DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum{
	"APPROVE": DelegatedResourceAccessRequestApprovalDetailsApprovalActionApprove,
	"REJECT":  DelegatedResourceAccessRequestApprovalDetailsApprovalActionReject,
}

var mappingDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnumLowerCase = map[string]DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum{
	"approve": DelegatedResourceAccessRequestApprovalDetailsApprovalActionApprove,
	"reject":  DelegatedResourceAccessRequestApprovalDetailsApprovalActionReject,
}

// GetDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnumValues Enumerates the set of values for DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum
func GetDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnumValues() []DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum {
	values := make([]DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum, 0)
	for _, v := range mappingDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnumStringValues Enumerates the set of values in String for DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum
func GetDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnumStringValues() []string {
	return []string{
		"APPROVE",
		"REJECT",
	}
}

// GetMappingDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum(val string) (DelegatedResourceAccessRequestApprovalDetailsApprovalActionEnum, bool) {
	enum, ok := mappingDelegatedResourceAccessRequestApprovalDetailsApprovalActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum Enum with underlying type: string
type DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum string

// Set of constants representing the allowable values for DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum
const (
	DelegatedResourceAccessRequestApprovalDetailsApprovalTypeInitial   DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum = "INITIAL"
	DelegatedResourceAccessRequestApprovalDetailsApprovalTypeExtension DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum = "EXTENSION"
)

var mappingDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum = map[string]DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum{
	"INITIAL":   DelegatedResourceAccessRequestApprovalDetailsApprovalTypeInitial,
	"EXTENSION": DelegatedResourceAccessRequestApprovalDetailsApprovalTypeExtension,
}

var mappingDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnumLowerCase = map[string]DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum{
	"initial":   DelegatedResourceAccessRequestApprovalDetailsApprovalTypeInitial,
	"extension": DelegatedResourceAccessRequestApprovalDetailsApprovalTypeExtension,
}

// GetDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnumValues Enumerates the set of values for DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum
func GetDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnumValues() []DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum {
	values := make([]DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum, 0)
	for _, v := range mappingDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnumStringValues Enumerates the set of values in String for DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum
func GetDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnumStringValues() []string {
	return []string{
		"INITIAL",
		"EXTENSION",
	}
}

// GetMappingDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum(val string) (DelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnum, bool) {
	enum, ok := mappingDelegatedResourceAccessRequestApprovalDetailsApprovalTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
