// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LockDurationDetails Details for setting a retention date or legal hold.
type LockDurationDetails struct {

	// The retention period (measured in days) defines how long a snapshot remains locked,
	// preventing user modifications or deletions. In governance mode this period can be adjusted,
	// but in compliance mode it becomes permanent after a cool-off period.
	// Snapshots can be locked for a minimum of 0 days and a maximum of 36,500 days.
	// A value of 0 days stands for an indefinite retention period and it is used for a legal hold.
	LockDuration *int `mandatory:"true" json:"lockDuration"`

	// Can be GOVERNANCE or COMPLIANCE.
	// GOVERNANCE MODE: locks snapshots based on either a retention period or a legal hold.
	// COMPLIANCE MODE: the customer can only remove the snapshot during its cooling-off period.
	// Once that time ends, the snapshot becomes immutable; customers cannot delete or modify it
	// until its set retention date passes. After the snapshot is locked, customers can only
	// increase its retention period.
	LockMode LockDurationDetailsLockModeEnum `mandatory:"true" json:"lockMode"`

	// For snapshots in compliance mode, a cooling-off period (measured in days) begins.
	// During this time, you can still edit or remove the lock. Once this period ends,
	// the snapshot becomes immutable until the specified retention date expires, permanently
	// preventing any deletion or modification. The cool off duration can be set for a minimum of
	// 0 days and a maximum of 365. It defaults to 14 days if not set.
	CoolOffDuration *int `mandatory:"false" json:"coolOffDuration"`
}

func (m LockDurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LockDurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLockDurationDetailsLockModeEnum(string(m.LockMode)); !ok && m.LockMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LockMode: %s. Supported values are: %s.", m.LockMode, strings.Join(GetLockDurationDetailsLockModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LockDurationDetailsLockModeEnum Enum with underlying type: string
type LockDurationDetailsLockModeEnum string

// Set of constants representing the allowable values for LockDurationDetailsLockModeEnum
const (
	LockDurationDetailsLockModeGovernance LockDurationDetailsLockModeEnum = "GOVERNANCE"
	LockDurationDetailsLockModeCompliance LockDurationDetailsLockModeEnum = "COMPLIANCE"
)

var mappingLockDurationDetailsLockModeEnum = map[string]LockDurationDetailsLockModeEnum{
	"GOVERNANCE": LockDurationDetailsLockModeGovernance,
	"COMPLIANCE": LockDurationDetailsLockModeCompliance,
}

var mappingLockDurationDetailsLockModeEnumLowerCase = map[string]LockDurationDetailsLockModeEnum{
	"governance": LockDurationDetailsLockModeGovernance,
	"compliance": LockDurationDetailsLockModeCompliance,
}

// GetLockDurationDetailsLockModeEnumValues Enumerates the set of values for LockDurationDetailsLockModeEnum
func GetLockDurationDetailsLockModeEnumValues() []LockDurationDetailsLockModeEnum {
	values := make([]LockDurationDetailsLockModeEnum, 0)
	for _, v := range mappingLockDurationDetailsLockModeEnum {
		values = append(values, v)
	}
	return values
}

// GetLockDurationDetailsLockModeEnumStringValues Enumerates the set of values in String for LockDurationDetailsLockModeEnum
func GetLockDurationDetailsLockModeEnumStringValues() []string {
	return []string{
		"GOVERNANCE",
		"COMPLIANCE",
	}
}

// GetMappingLockDurationDetailsLockModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLockDurationDetailsLockModeEnum(val string) (LockDurationDetailsLockModeEnum, bool) {
	enum, ok := mappingLockDurationDetailsLockModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
