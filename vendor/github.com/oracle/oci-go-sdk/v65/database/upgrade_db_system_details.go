// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpgradeDbSystemDetails Details for upgrading the operating system and Oracle Grid Infrastructure (GI) of a DB system.
type UpgradeDbSystemDetails struct {

	// The operating system upgrade action.
	Action UpgradeDbSystemDetailsActionEnum `mandatory:"true" json:"action"`

	// The retention period, in days, for the snapshot that allows you to perform a rollback of the upgrade operation. After this number of days passes, you cannot roll back the upgrade.
	SnapshotRetentionPeriodInDays *int `mandatory:"false" json:"snapshotRetentionPeriodInDays"`

	// A valid Oracle Grid Infrastructure (GI) software version.
	NewGiVersion *string `mandatory:"false" json:"newGiVersion"`

	// If true, rollback time is updated even if operating system upgrade history contains errors.
	IsSnapshotRetentionDaysForceUpdated *bool `mandatory:"false" json:"isSnapshotRetentionDaysForceUpdated"`
}

func (m UpgradeDbSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpgradeDbSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUpgradeDbSystemDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetUpgradeDbSystemDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpgradeDbSystemDetailsActionEnum Enum with underlying type: string
type UpgradeDbSystemDetailsActionEnum string

// Set of constants representing the allowable values for UpgradeDbSystemDetailsActionEnum
const (
	UpgradeDbSystemDetailsActionPrecheck                    UpgradeDbSystemDetailsActionEnum = "PRECHECK"
	UpgradeDbSystemDetailsActionRollback                    UpgradeDbSystemDetailsActionEnum = "ROLLBACK"
	UpgradeDbSystemDetailsActionUpdateSnapshotRetentionDays UpgradeDbSystemDetailsActionEnum = "UPDATE_SNAPSHOT_RETENTION_DAYS"
	UpgradeDbSystemDetailsActionUpgrade                     UpgradeDbSystemDetailsActionEnum = "UPGRADE"
)

var mappingUpgradeDbSystemDetailsActionEnum = map[string]UpgradeDbSystemDetailsActionEnum{
	"PRECHECK":                       UpgradeDbSystemDetailsActionPrecheck,
	"ROLLBACK":                       UpgradeDbSystemDetailsActionRollback,
	"UPDATE_SNAPSHOT_RETENTION_DAYS": UpgradeDbSystemDetailsActionUpdateSnapshotRetentionDays,
	"UPGRADE":                        UpgradeDbSystemDetailsActionUpgrade,
}

var mappingUpgradeDbSystemDetailsActionEnumLowerCase = map[string]UpgradeDbSystemDetailsActionEnum{
	"precheck":                       UpgradeDbSystemDetailsActionPrecheck,
	"rollback":                       UpgradeDbSystemDetailsActionRollback,
	"update_snapshot_retention_days": UpgradeDbSystemDetailsActionUpdateSnapshotRetentionDays,
	"upgrade":                        UpgradeDbSystemDetailsActionUpgrade,
}

// GetUpgradeDbSystemDetailsActionEnumValues Enumerates the set of values for UpgradeDbSystemDetailsActionEnum
func GetUpgradeDbSystemDetailsActionEnumValues() []UpgradeDbSystemDetailsActionEnum {
	values := make([]UpgradeDbSystemDetailsActionEnum, 0)
	for _, v := range mappingUpgradeDbSystemDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUpgradeDbSystemDetailsActionEnumStringValues Enumerates the set of values in String for UpgradeDbSystemDetailsActionEnum
func GetUpgradeDbSystemDetailsActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"ROLLBACK",
		"UPDATE_SNAPSHOT_RETENTION_DAYS",
		"UPGRADE",
	}
}

// GetMappingUpgradeDbSystemDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpgradeDbSystemDetailsActionEnum(val string) (UpgradeDbSystemDetailsActionEnum, bool) {
	enum, ok := mappingUpgradeDbSystemDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
