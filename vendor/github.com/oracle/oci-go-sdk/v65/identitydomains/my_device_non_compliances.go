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

// MyDeviceNonCompliances Device Non Compliances
type MyDeviceNonCompliances struct {

	// Device Compliance name
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Name *string `mandatory:"true" json:"name"`

	// Device Compliance value
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Value *string `mandatory:"true" json:"value"`

	// Device Compliance Action
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	//  - idcsRequiresWriteForAccessFlows: true
	//  - idcsRequiresImmediateReadAfterWriteForAccessFlows: true
	Action MyDeviceNonCompliancesActionEnum `mandatory:"true" json:"action"`
}

func (m MyDeviceNonCompliances) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MyDeviceNonCompliances) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMyDeviceNonCompliancesActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetMyDeviceNonCompliancesActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MyDeviceNonCompliancesActionEnum Enum with underlying type: string
type MyDeviceNonCompliancesActionEnum string

// Set of constants representing the allowable values for MyDeviceNonCompliancesActionEnum
const (
	MyDeviceNonCompliancesActionNotify  MyDeviceNonCompliancesActionEnum = "NOTIFY"
	MyDeviceNonCompliancesActionBlock   MyDeviceNonCompliancesActionEnum = "BLOCK"
	MyDeviceNonCompliancesActionAllow   MyDeviceNonCompliancesActionEnum = "ALLOW"
	MyDeviceNonCompliancesActionUnknown MyDeviceNonCompliancesActionEnum = "UNKNOWN"
)

var mappingMyDeviceNonCompliancesActionEnum = map[string]MyDeviceNonCompliancesActionEnum{
	"NOTIFY":  MyDeviceNonCompliancesActionNotify,
	"BLOCK":   MyDeviceNonCompliancesActionBlock,
	"ALLOW":   MyDeviceNonCompliancesActionAllow,
	"UNKNOWN": MyDeviceNonCompliancesActionUnknown,
}

var mappingMyDeviceNonCompliancesActionEnumLowerCase = map[string]MyDeviceNonCompliancesActionEnum{
	"notify":  MyDeviceNonCompliancesActionNotify,
	"block":   MyDeviceNonCompliancesActionBlock,
	"allow":   MyDeviceNonCompliancesActionAllow,
	"unknown": MyDeviceNonCompliancesActionUnknown,
}

// GetMyDeviceNonCompliancesActionEnumValues Enumerates the set of values for MyDeviceNonCompliancesActionEnum
func GetMyDeviceNonCompliancesActionEnumValues() []MyDeviceNonCompliancesActionEnum {
	values := make([]MyDeviceNonCompliancesActionEnum, 0)
	for _, v := range mappingMyDeviceNonCompliancesActionEnum {
		values = append(values, v)
	}
	return values
}

// GetMyDeviceNonCompliancesActionEnumStringValues Enumerates the set of values in String for MyDeviceNonCompliancesActionEnum
func GetMyDeviceNonCompliancesActionEnumStringValues() []string {
	return []string{
		"NOTIFY",
		"BLOCK",
		"ALLOW",
		"UNKNOWN",
	}
}

// GetMappingMyDeviceNonCompliancesActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMyDeviceNonCompliancesActionEnum(val string) (MyDeviceNonCompliancesActionEnum, bool) {
	enum, ok := mappingMyDeviceNonCompliancesActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
