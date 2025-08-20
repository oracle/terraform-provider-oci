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

// UpdateBackupScheduleDetails Defines the backup schedule details for update operation.
type UpdateBackupScheduleDetails struct {

	// The start timestamp for the deployment backup schedule. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-10-25T18:19:29.600Z`.
	TimeBackupScheduled *common.SDKTime `mandatory:"false" json:"timeBackupScheduled"`

	// The frequency of the deployment backup schedule. Frequency can be DAILY, WEEKLY or MONTHLY.
	FrequencyBackupScheduled UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum `mandatory:"false" json:"frequencyBackupScheduled,omitempty"`

	// Name of the bucket where the object is to be uploaded in the object storage
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Name of namespace that serves as a container for all of your buckets
	NamespaceName *string `mandatory:"false" json:"namespaceName"`

	// Parameter to allow users to create backup without trails
	IsMetadataOnly *bool `mandatory:"false" json:"isMetadataOnly"`
}

func (m UpdateBackupScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateBackupScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateBackupScheduleDetailsFrequencyBackupScheduledEnum(string(m.FrequencyBackupScheduled)); !ok && m.FrequencyBackupScheduled != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FrequencyBackupScheduled: %s. Supported values are: %s.", m.FrequencyBackupScheduled, strings.Join(GetUpdateBackupScheduleDetailsFrequencyBackupScheduledEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum Enum with underlying type: string
type UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum string

// Set of constants representing the allowable values for UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum
const (
	UpdateBackupScheduleDetailsFrequencyBackupScheduledDaily   UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum = "DAILY"
	UpdateBackupScheduleDetailsFrequencyBackupScheduledWeekly  UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum = "WEEKLY"
	UpdateBackupScheduleDetailsFrequencyBackupScheduledMonthly UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum = "MONTHLY"
)

var mappingUpdateBackupScheduleDetailsFrequencyBackupScheduledEnum = map[string]UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum{
	"DAILY":   UpdateBackupScheduleDetailsFrequencyBackupScheduledDaily,
	"WEEKLY":  UpdateBackupScheduleDetailsFrequencyBackupScheduledWeekly,
	"MONTHLY": UpdateBackupScheduleDetailsFrequencyBackupScheduledMonthly,
}

var mappingUpdateBackupScheduleDetailsFrequencyBackupScheduledEnumLowerCase = map[string]UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum{
	"daily":   UpdateBackupScheduleDetailsFrequencyBackupScheduledDaily,
	"weekly":  UpdateBackupScheduleDetailsFrequencyBackupScheduledWeekly,
	"monthly": UpdateBackupScheduleDetailsFrequencyBackupScheduledMonthly,
}

// GetUpdateBackupScheduleDetailsFrequencyBackupScheduledEnumValues Enumerates the set of values for UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum
func GetUpdateBackupScheduleDetailsFrequencyBackupScheduledEnumValues() []UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum {
	values := make([]UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum, 0)
	for _, v := range mappingUpdateBackupScheduleDetailsFrequencyBackupScheduledEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateBackupScheduleDetailsFrequencyBackupScheduledEnumStringValues Enumerates the set of values in String for UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum
func GetUpdateBackupScheduleDetailsFrequencyBackupScheduledEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
}

// GetMappingUpdateBackupScheduleDetailsFrequencyBackupScheduledEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateBackupScheduleDetailsFrequencyBackupScheduledEnum(val string) (UpdateBackupScheduleDetailsFrequencyBackupScheduledEnum, bool) {
	enum, ok := mappingUpdateBackupScheduleDetailsFrequencyBackupScheduledEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
