// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ChangeDisasterRecoveryConfigurationDetails Details to update the cross-region disaster recovery (DR) details of the standby Autonomous Database Serverless instance.
type ChangeDisasterRecoveryConfigurationDetails struct {

	// Indicates the disaster recovery (DR) type of the Autonomous Database Serverless instance.
	// Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover.
	// Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover.
	DisasterRecoveryType ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum `mandatory:"false" json:"disasterRecoveryType,omitempty"`

	// Time and date stored as an RFC 3339 formatted timestamp string. For example, 2022-01-01T12:00:00.000Z would set a limit for the snapshot standby to be converted back to a cross-region standby database.
	TimeSnapshotStandbyEnabledTill *common.SDKTime `mandatory:"false" json:"timeSnapshotStandbyEnabledTill"`

	// Indicates if user wants to convert to a snapshot standby. For example, true would set a standby database to snapshot standby database. False would set a snapshot standby database back to regular standby database.
	IsSnapshotStandby *bool `mandatory:"false" json:"isSnapshotStandby"`

	// If true, 7 days worth of backups are replicated across regions for Cross-Region ADB or Backup-Based DR between Primary and Standby. If false, the backups taken on the Primary are not replicated to the Standby database.
	IsReplicateAutomaticBackups *bool `mandatory:"false" json:"isReplicateAutomaticBackups"`
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
	ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeAdg         ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum = "ADG"
	ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeBackupBased ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum = "BACKUP_BASED"
)

var mappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum = map[string]ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum{
	"ADG":          ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeAdg,
	"BACKUP_BASED": ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeBackupBased,
}

var mappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumLowerCase = map[string]ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum{
	"adg":          ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeAdg,
	"backup_based": ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeBackupBased,
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
	}
}

// GetMappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum(val string) (ChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnum, bool) {
	enum, ok := mappingChangeDisasterRecoveryConfigurationDetailsDisasterRecoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
