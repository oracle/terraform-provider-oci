// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BackupSchedule Defines the schedule of the deployment backup.
type BackupSchedule struct {

	// The start timestamp for the deployment backup schedule. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-10-25T18:19:29.600Z`.
	TimeBackupScheduled *common.SDKTime `mandatory:"true" json:"timeBackupScheduled"`

	// The frequency of the deployment backup schedule. Frequency can be DAILY, WEEKLY or MONTHLY.
	FrequencyBackupScheduled BackupScheduleFrequencyBackupScheduledEnum `mandatory:"true" json:"frequencyBackupScheduled"`

	// Name of the bucket where the object is to be uploaded in the object storage
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of namespace that serves as a container for all of your buckets
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// Parameter to allow users to create backup without trails
	IsMetadataOnly *bool `mandatory:"true" json:"isMetadataOnly"`
}

func (m BackupSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BackupSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBackupScheduleFrequencyBackupScheduledEnum(string(m.FrequencyBackupScheduled)); !ok && m.FrequencyBackupScheduled != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FrequencyBackupScheduled: %s. Supported values are: %s.", m.FrequencyBackupScheduled, strings.Join(GetBackupScheduleFrequencyBackupScheduledEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BackupScheduleFrequencyBackupScheduledEnum Enum with underlying type: string
type BackupScheduleFrequencyBackupScheduledEnum string

// Set of constants representing the allowable values for BackupScheduleFrequencyBackupScheduledEnum
const (
	BackupScheduleFrequencyBackupScheduledDaily   BackupScheduleFrequencyBackupScheduledEnum = "DAILY"
	BackupScheduleFrequencyBackupScheduledWeekly  BackupScheduleFrequencyBackupScheduledEnum = "WEEKLY"
	BackupScheduleFrequencyBackupScheduledMonthly BackupScheduleFrequencyBackupScheduledEnum = "MONTHLY"
)

var mappingBackupScheduleFrequencyBackupScheduledEnum = map[string]BackupScheduleFrequencyBackupScheduledEnum{
	"DAILY":   BackupScheduleFrequencyBackupScheduledDaily,
	"WEEKLY":  BackupScheduleFrequencyBackupScheduledWeekly,
	"MONTHLY": BackupScheduleFrequencyBackupScheduledMonthly,
}

var mappingBackupScheduleFrequencyBackupScheduledEnumLowerCase = map[string]BackupScheduleFrequencyBackupScheduledEnum{
	"daily":   BackupScheduleFrequencyBackupScheduledDaily,
	"weekly":  BackupScheduleFrequencyBackupScheduledWeekly,
	"monthly": BackupScheduleFrequencyBackupScheduledMonthly,
}

// GetBackupScheduleFrequencyBackupScheduledEnumValues Enumerates the set of values for BackupScheduleFrequencyBackupScheduledEnum
func GetBackupScheduleFrequencyBackupScheduledEnumValues() []BackupScheduleFrequencyBackupScheduledEnum {
	values := make([]BackupScheduleFrequencyBackupScheduledEnum, 0)
	for _, v := range mappingBackupScheduleFrequencyBackupScheduledEnum {
		values = append(values, v)
	}
	return values
}

// GetBackupScheduleFrequencyBackupScheduledEnumStringValues Enumerates the set of values in String for BackupScheduleFrequencyBackupScheduledEnum
func GetBackupScheduleFrequencyBackupScheduledEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
}

// GetMappingBackupScheduleFrequencyBackupScheduledEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBackupScheduleFrequencyBackupScheduledEnum(val string) (BackupScheduleFrequencyBackupScheduledEnum, bool) {
	enum, ok := mappingBackupScheduleFrequencyBackupScheduledEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
