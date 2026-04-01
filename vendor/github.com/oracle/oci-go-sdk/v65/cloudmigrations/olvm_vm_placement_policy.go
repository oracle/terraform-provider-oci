// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmVmPlacementPolicy The configuration of the virtual machine???s placement policy.
type OlvmVmPlacementPolicy struct {

	// Affinity of the virtual machine.
	VmAffinity OlvmVmPlacementPolicyVmAffinityEnum `mandatory:"false" json:"vmAffinity,omitempty"`

	// List of hosts.
	Hosts []OlvmHost `mandatory:"false" json:"hosts"`
}

func (m OlvmVmPlacementPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmVmPlacementPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmVmPlacementPolicyVmAffinityEnum(string(m.VmAffinity)); !ok && m.VmAffinity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VmAffinity: %s. Supported values are: %s.", m.VmAffinity, strings.Join(GetOlvmVmPlacementPolicyVmAffinityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmVmPlacementPolicyVmAffinityEnum Enum with underlying type: string
type OlvmVmPlacementPolicyVmAffinityEnum string

// Set of constants representing the allowable values for OlvmVmPlacementPolicyVmAffinityEnum
const (
	OlvmVmPlacementPolicyVmAffinityMigratable     OlvmVmPlacementPolicyVmAffinityEnum = "MIGRATABLE"
	OlvmVmPlacementPolicyVmAffinityPinned         OlvmVmPlacementPolicyVmAffinityEnum = "PINNED"
	OlvmVmPlacementPolicyVmAffinityUserMigratable OlvmVmPlacementPolicyVmAffinityEnum = "USER_MIGRATABLE"
)

var mappingOlvmVmPlacementPolicyVmAffinityEnum = map[string]OlvmVmPlacementPolicyVmAffinityEnum{
	"MIGRATABLE":      OlvmVmPlacementPolicyVmAffinityMigratable,
	"PINNED":          OlvmVmPlacementPolicyVmAffinityPinned,
	"USER_MIGRATABLE": OlvmVmPlacementPolicyVmAffinityUserMigratable,
}

var mappingOlvmVmPlacementPolicyVmAffinityEnumLowerCase = map[string]OlvmVmPlacementPolicyVmAffinityEnum{
	"migratable":      OlvmVmPlacementPolicyVmAffinityMigratable,
	"pinned":          OlvmVmPlacementPolicyVmAffinityPinned,
	"user_migratable": OlvmVmPlacementPolicyVmAffinityUserMigratable,
}

// GetOlvmVmPlacementPolicyVmAffinityEnumValues Enumerates the set of values for OlvmVmPlacementPolicyVmAffinityEnum
func GetOlvmVmPlacementPolicyVmAffinityEnumValues() []OlvmVmPlacementPolicyVmAffinityEnum {
	values := make([]OlvmVmPlacementPolicyVmAffinityEnum, 0)
	for _, v := range mappingOlvmVmPlacementPolicyVmAffinityEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmVmPlacementPolicyVmAffinityEnumStringValues Enumerates the set of values in String for OlvmVmPlacementPolicyVmAffinityEnum
func GetOlvmVmPlacementPolicyVmAffinityEnumStringValues() []string {
	return []string{
		"MIGRATABLE",
		"PINNED",
		"USER_MIGRATABLE",
	}
}

// GetMappingOlvmVmPlacementPolicyVmAffinityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmVmPlacementPolicyVmAffinityEnum(val string) (OlvmVmPlacementPolicyVmAffinityEnum, bool) {
	enum, ok := mappingOlvmVmPlacementPolicyVmAffinityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
