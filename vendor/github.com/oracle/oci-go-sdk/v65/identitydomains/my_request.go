// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MyRequest Request resource
type MyRequest struct {

	// REQUIRED. The schemas attribute is an array of Strings which allows introspection of the supported schema version for a SCIM representation as well any schema extensions supported by that representation. Each String value must be a unique URI. This specification defines URIs for User, Group, and a standard \"enterprise\" extension. All representations of SCIM schema MUST include a non-zero value array with value(s) of the URIs supported by that representation. Duplicate values MUST NOT be included. Value order is not specified and MUST not impact behavior.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Schemas []string `mandatory:"true" json:"schemas"`

	// justification
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Justification *string `mandatory:"true" json:"justification"`

	Requesting *MyRequestRequesting `mandatory:"true" json:"requesting"`

	// Unique identifier for the SCIM Resource as defined by the Service Provider. Each representation of the Resource MUST include a non-empty id value. This identifier MUST be unique across the Service Provider's entire set of Resources. It MUST be a stable, non-reassignable identifier that does not change when the same Resource is returned in subsequent requests. The value of the id attribute is always issued by the Service Provider and MUST never be specified by the Service Consumer. bulkId: is a reserved keyword and MUST NOT be used in the unique identifier.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: always
	//  - type: string
	//  - uniqueness: global
	Id *string `mandatory:"false" json:"id"`

	// Unique OCI identifier for the SCIM Resource.
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: immutable
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: global
	Ocid *string `mandatory:"false" json:"ocid"`

	Meta *Meta `mandatory:"false" json:"meta"`

	IdcsCreatedBy *IdcsCreatedBy `mandatory:"false" json:"idcsCreatedBy"`

	IdcsLastModifiedBy *IdcsLastModifiedBy `mandatory:"false" json:"idcsLastModifiedBy"`

	// Each value of this attribute specifies an operation that only an internal client may perform on this particular resource.
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsPreventedOperations []IdcsPreventedOperationsEnum `mandatory:"false" json:"idcsPreventedOperations,omitempty"`

	// A list of tags on this resource.
	// **SCIM++ Properties:**
	//  - idcsCompositeKey: [key, value]
	//  - idcsSearchable: true
	//  - multiValued: true
	//  - mutability: readWrite
	//  - required: false
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	Tags []Tags `mandatory:"false" json:"tags"`

	// A boolean flag indicating this resource in the process of being deleted. Usually set to true when synchronous deletion of the resource would take too long.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	DeleteInProgress *bool `mandatory:"false" json:"deleteInProgress"`

	// The release number when the resource was upgraded.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: request
	//  - type: string
	//  - uniqueness: none
	IdcsLastUpgradedInRelease *string `mandatory:"false" json:"idcsLastUpgradedInRelease"`

	// OCI Domain Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DomainOcid *string `mandatory:"false" json:"domainOcid"`

	// OCI Compartment Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	CompartmentOcid *string `mandatory:"false" json:"compartmentOcid"`

	// OCI Tenant Id (ocid) in which the resource lives.
	// **SCIM++ Properties:**
	//  - caseExact: false
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	TenancyOcid *string `mandatory:"false" json:"tenancyOcid"`

	// status
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Status MyRequestStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Requestor can set action to CANCEL to cancel the request or to ESCALATE to escalate the request while the request status is IN_PROGRESS. Requestor can't escalate the request if canceling or escalation is in progress.
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: false
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	Action MyRequestActionEnum `mandatory:"false" json:"action,omitempty"`

	// Time by when Request expires
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: false
	//  - returned: default
	//  - type: dateTime
	//  - uniqueness: none
	Expires *string `mandatory:"false" json:"expires"`

	// Approvals created for this request.
	// **Added In:** 2307071836
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: true
	//  - mutability: readOnly
	//  - returned: request
	//  - type: complex
	//  - uniqueness: none
	ApprovalDetails []MyRequestApprovalDetails `mandatory:"false" json:"approvalDetails"`

	Requestor *MyRequestRequestor `mandatory:"false" json:"requestor"`
}

func (m MyRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.IdcsPreventedOperations {
		if _, ok := GetMappingIdcsPreventedOperationsEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IdcsPreventedOperations: %s. Supported values are: %s.", val, strings.Join(GetIdcsPreventedOperationsEnumStringValues(), ",")))
		}
	}

	if _, ok := GetMappingMyRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetMyRequestStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMyRequestActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetMyRequestActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyRequestStatusEnum Enum with underlying type: string
type MyRequestStatusEnum string

// Set of constants representing the allowable values for MyRequestStatusEnum
const (
	MyRequestStatusCreated    MyRequestStatusEnum = "CREATED"
	MyRequestStatusComplete   MyRequestStatusEnum = "COMPLETE"
	MyRequestStatusInProgress MyRequestStatusEnum = "IN_PROGRESS"
	MyRequestStatusApproved   MyRequestStatusEnum = "APPROVED"
	MyRequestStatusRejected   MyRequestStatusEnum = "REJECTED"
	MyRequestStatusCanceled   MyRequestStatusEnum = "CANCELED"
	MyRequestStatusExpired    MyRequestStatusEnum = "EXPIRED"
	MyRequestStatusFailed     MyRequestStatusEnum = "FAILED"
)

var mappingMyRequestStatusEnum = map[string]MyRequestStatusEnum{
	"CREATED":     MyRequestStatusCreated,
	"COMPLETE":    MyRequestStatusComplete,
	"IN_PROGRESS": MyRequestStatusInProgress,
	"APPROVED":    MyRequestStatusApproved,
	"REJECTED":    MyRequestStatusRejected,
	"CANCELED":    MyRequestStatusCanceled,
	"EXPIRED":     MyRequestStatusExpired,
	"FAILED":      MyRequestStatusFailed,
}

var mappingMyRequestStatusEnumLowerCase = map[string]MyRequestStatusEnum{
	"created":     MyRequestStatusCreated,
	"complete":    MyRequestStatusComplete,
	"in_progress": MyRequestStatusInProgress,
	"approved":    MyRequestStatusApproved,
	"rejected":    MyRequestStatusRejected,
	"canceled":    MyRequestStatusCanceled,
	"expired":     MyRequestStatusExpired,
	"failed":      MyRequestStatusFailed,
}

// GetMyRequestStatusEnumValues Enumerates the set of values for MyRequestStatusEnum
func GetMyRequestStatusEnumValues() []MyRequestStatusEnum {
	values := make([]MyRequestStatusEnum, 0)
	for _, v := range mappingMyRequestStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMyRequestStatusEnumStringValues Enumerates the set of values in String for MyRequestStatusEnum
func GetMyRequestStatusEnumStringValues() []string {
	return []string{
		"CREATED",
		"COMPLETE",
		"IN_PROGRESS",
		"APPROVED",
		"REJECTED",
		"CANCELED",
		"EXPIRED",
		"FAILED",
	}
}

// GetMappingMyRequestStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyRequestStatusEnum(val string) (MyRequestStatusEnum, bool) {
	enum, ok := mappingMyRequestStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MyRequestActionEnum Enum with underlying type: string
type MyRequestActionEnum string

// Set of constants representing the allowable values for MyRequestActionEnum
const (
	MyRequestActionCancel   MyRequestActionEnum = "CANCEL"
	MyRequestActionEscalate MyRequestActionEnum = "ESCALATE"
)

var mappingMyRequestActionEnum = map[string]MyRequestActionEnum{
	"CANCEL":   MyRequestActionCancel,
	"ESCALATE": MyRequestActionEscalate,
}

var mappingMyRequestActionEnumLowerCase = map[string]MyRequestActionEnum{
	"cancel":   MyRequestActionCancel,
	"escalate": MyRequestActionEscalate,
}

// GetMyRequestActionEnumValues Enumerates the set of values for MyRequestActionEnum
func GetMyRequestActionEnumValues() []MyRequestActionEnum {
	values := make([]MyRequestActionEnum, 0)
	for _, v := range mappingMyRequestActionEnum {
		values = append(values, v)
	}
	return values
}

// GetMyRequestActionEnumStringValues Enumerates the set of values in String for MyRequestActionEnum
func GetMyRequestActionEnumStringValues() []string {
	return []string{
		"CANCEL",
		"ESCALATE",
	}
}

// GetMappingMyRequestActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyRequestActionEnum(val string) (MyRequestActionEnum, bool) {
	enum, ok := mappingMyRequestActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
