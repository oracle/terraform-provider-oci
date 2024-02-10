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

// DisasterRecoveryConfiguration Configurations of a Disaster Recovery.
type DisasterRecoveryConfiguration struct {

	// Indicates the disaster recovery (DR) type of the Autonomous Database Serverless instance.
	// Autonomous Data Guard (ADG) DR type provides business critical DR with a faster recovery time objective (RTO) during failover or switchover.
	// Backup-based DR type provides lower cost DR with a slower RTO during failover or switchover.
	DisasterRecoveryType DisasterRecoveryConfigurationDisasterRecoveryTypeEnum `mandatory:"false" json:"disasterRecoveryType,omitempty"`

	// Time and date stored as an RFC 3339 formatted timestamp string. For example, 2022-01-01T12:00:00.000Z would set a limit for the snapshot standby to be converted back to a cross-region standby database.
	TimeSnapshotStandbyEnabledTill *common.SDKTime `mandatory:"false" json:"timeSnapshotStandbyEnabledTill"`

	// Indicates if user wants to convert to a snapshot standby. For example, true would set a standby database to snapshot standby database. False would set a snapshot standby database back to regular standby database.
	IsSnapshotStandby *bool `mandatory:"false" json:"isSnapshotStandby"`

	// If true, 7 days worth of backups are replicated across regions for Cross-Region ADB or Backup-Based DR between Primary and Standby. If false, the backups taken on the Primary are not replicated to the Standby database.
	IsReplicateAutomaticBackups *bool `mandatory:"false" json:"isReplicateAutomaticBackups"`
}

func (m DisasterRecoveryConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DisasterRecoveryConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum(string(m.DisasterRecoveryType)); !ok && m.DisasterRecoveryType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DisasterRecoveryType: %s. Supported values are: %s.", m.DisasterRecoveryType, strings.Join(GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DisasterRecoveryConfigurationDisasterRecoveryTypeEnum Enum with underlying type: string
type DisasterRecoveryConfigurationDisasterRecoveryTypeEnum string

// Set of constants representing the allowable values for DisasterRecoveryConfigurationDisasterRecoveryTypeEnum
const (
	DisasterRecoveryConfigurationDisasterRecoveryTypeAdg         DisasterRecoveryConfigurationDisasterRecoveryTypeEnum = "ADG"
	DisasterRecoveryConfigurationDisasterRecoveryTypeBackupBased DisasterRecoveryConfigurationDisasterRecoveryTypeEnum = "BACKUP_BASED"
)

var mappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum = map[string]DisasterRecoveryConfigurationDisasterRecoveryTypeEnum{
	"ADG":          DisasterRecoveryConfigurationDisasterRecoveryTypeAdg,
	"BACKUP_BASED": DisasterRecoveryConfigurationDisasterRecoveryTypeBackupBased,
}

var mappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnumLowerCase = map[string]DisasterRecoveryConfigurationDisasterRecoveryTypeEnum{
	"adg":          DisasterRecoveryConfigurationDisasterRecoveryTypeAdg,
	"backup_based": DisasterRecoveryConfigurationDisasterRecoveryTypeBackupBased,
}

// GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumValues Enumerates the set of values for DisasterRecoveryConfigurationDisasterRecoveryTypeEnum
func GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumValues() []DisasterRecoveryConfigurationDisasterRecoveryTypeEnum {
	values := make([]DisasterRecoveryConfigurationDisasterRecoveryTypeEnum, 0)
	for _, v := range mappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumStringValues Enumerates the set of values in String for DisasterRecoveryConfigurationDisasterRecoveryTypeEnum
func GetDisasterRecoveryConfigurationDisasterRecoveryTypeEnumStringValues() []string {
	return []string{
		"ADG",
		"BACKUP_BASED",
	}
}

// GetMappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnum(val string) (DisasterRecoveryConfigurationDisasterRecoveryTypeEnum, bool) {
	enum, ok := mappingDisasterRecoveryConfigurationDisasterRecoveryTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
