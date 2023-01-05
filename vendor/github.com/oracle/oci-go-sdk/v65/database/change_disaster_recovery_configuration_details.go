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

// ChangeDisasterRecoveryConfigurationDetails Details to update the cross-region Disaster Recovery (DR) details of the Standby Autonomous Database on shared Exadata infrastructure.
type ChangeDisasterRecoveryConfigurationDetails struct {

	// Indicates the disaster recovery (DR) type of the Shared Autonomous Database.
	// Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover.
	// Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover.
	DisasterRecoveryType ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum `mandatory:"false" json:"disasterRecoveryType,omitempty"`

	// Time and date stored as an RFC 3339 formatted timestamp string. For example, 2022-01-01T12:00:00.000Z would set a limit for the snapshot standby to be converted back to a cross-region standby database.
	TimeSnapshotStandbyEnabledTill *common.SDKTime `mandatory:"false" json:"timeSnapshotStandbyEnabledTill"`
}

func (m ChangeDisasterRecoveryConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChangeDisasterRecoveryConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum(string(m.DisasterRecoveryType)); !ok && m.DisasterRecoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryType: %s. Supported values are: %s.", m.DisasterRecoveryType, strings.Join(GetChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum Enum with underlying type: string
type ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum string

// Set of constants representing the allowable values for ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum
const (
	ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeAdg             ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum = "ADG"
	ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeBackupBased     ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum = "BACKUP_BASED"
	ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeSnapshotStandby ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum = "SNAPSHOT_STANDBY"
)

var mappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum = map[string]ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum{
	"ADG":              ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeAdg,
	"BACKUP_BASED":     ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeBackupBased,
	"SNAPSHOT_STANDBY": ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeSnapshotStandby,
}

var mappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumLowerCase = map[string]ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum{
	"adg":              ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeAdg,
	"backup_based":     ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeBackupBased,
	"snapshot_standby": ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeSnapshotStandby,
}

// GetChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumValues Enumerates the set of values for ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum
func GetChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumValues() []ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum {
	values := make([]ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum, 0)
	for _, v := range mappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumStringValues Enumerates the set of values in String for ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum
func GetChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumStringValues() []string {
	return []string{
		"ADG",
		"BACKUP_BASED",
		"SNAPSHOT_STANDBY",
	}
}

// GetMappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum(val string) (ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum, bool) {
	enum, ok := mappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
