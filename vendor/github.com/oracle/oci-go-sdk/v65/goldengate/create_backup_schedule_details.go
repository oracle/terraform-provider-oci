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

// CreateBackupScheduleDetails Defines the backup schedule details for create operation.
type CreateBackupScheduleDetails struct {

	// The start timestamp for the deployment backup schedule. The format is defined by
	// RFC3339 (https://tools.ietf.org/html/rfc3339), such as `2024-10-25T18:19:29.600Z`.
	TimeBackupScheduled *common.SDKTime `mandatory:"true" json:"timeBackupScheduled"`

	// The frequency of the deployment backup schedule. Frequency can be DAILY, WEEKLY or MONTHLY.
	FrequencyBackupScheduled CreateBackupScheduleDetailsFrequencyBackupScheduledEnum `mandatory:"true" json:"frequencyBackupScheduled"`

	// Name of the bucket where the object is to be uploaded in the object storage
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment being referenced.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of namespace that serves as a container for all of your buckets
	NamespaceName *string `mandatory:"true" json:"namespaceName"`

	// Parameter to allow users to create backup without trails
	IsMetadataOnly *bool `mandatory:"true" json:"isMetadataOnly"`
}

func (m CreateBackupScheduleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBackupScheduleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateBackupScheduleDetailsFrequencyBackupScheduledEnum(string(m.FrequencyBackupScheduled)); !ok && m.FrequencyBackupScheduled != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FrequencyBackupScheduled: %s. Supported values are: %s.", m.FrequencyBackupScheduled, strings.Join(GetCreateBackupScheduleDetailsFrequencyBackupScheduledEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateBackupScheduleDetailsFrequencyBackupScheduledEnum Enum with underlying type: string
type CreateBackupScheduleDetailsFrequencyBackupScheduledEnum string

// Set of constants representing the allowable values for CreateBackupScheduleDetailsFrequencyBackupScheduledEnum
const (
	CreateBackupScheduleDetailsFrequencyBackupScheduledDaily   CreateBackupScheduleDetailsFrequencyBackupScheduledEnum = "DAILY"
	CreateBackupScheduleDetailsFrequencyBackupScheduledWeekly  CreateBackupScheduleDetailsFrequencyBackupScheduledEnum = "WEEKLY"
	CreateBackupScheduleDetailsFrequencyBackupScheduledMonthly CreateBackupScheduleDetailsFrequencyBackupScheduledEnum = "MONTHLY"
)

var mappingCreateBackupScheduleDetailsFrequencyBackupScheduledEnum = map[string]CreateBackupScheduleDetailsFrequencyBackupScheduledEnum{
	"DAILY":   CreateBackupScheduleDetailsFrequencyBackupScheduledDaily,
	"WEEKLY":  CreateBackupScheduleDetailsFrequencyBackupScheduledWeekly,
	"MONTHLY": CreateBackupScheduleDetailsFrequencyBackupScheduledMonthly,
}

var mappingCreateBackupScheduleDetailsFrequencyBackupScheduledEnumLowerCase = map[string]CreateBackupScheduleDetailsFrequencyBackupScheduledEnum{
	"daily":   CreateBackupScheduleDetailsFrequencyBackupScheduledDaily,
	"weekly":  CreateBackupScheduleDetailsFrequencyBackupScheduledWeekly,
	"monthly": CreateBackupScheduleDetailsFrequencyBackupScheduledMonthly,
}

// GetCreateBackupScheduleDetailsFrequencyBackupScheduledEnumValues Enumerates the set of values for CreateBackupScheduleDetailsFrequencyBackupScheduledEnum
func GetCreateBackupScheduleDetailsFrequencyBackupScheduledEnumValues() []CreateBackupScheduleDetailsFrequencyBackupScheduledEnum {
	values := make([]CreateBackupScheduleDetailsFrequencyBackupScheduledEnum, 0)
	for _, v := range mappingCreateBackupScheduleDetailsFrequencyBackupScheduledEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBackupScheduleDetailsFrequencyBackupScheduledEnumStringValues Enumerates the set of values in String for CreateBackupScheduleDetailsFrequencyBackupScheduledEnum
func GetCreateBackupScheduleDetailsFrequencyBackupScheduledEnumStringValues() []string {
	return []string{
		"DAILY",
		"WEEKLY",
		"MONTHLY",
	}
}

// GetMappingCreateBackupScheduleDetailsFrequencyBackupScheduledEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBackupScheduleDetailsFrequencyBackupScheduledEnum(val string) (CreateBackupScheduleDetailsFrequencyBackupScheduledEnum, bool) {
	enum, ok := mappingCreateBackupScheduleDetailsFrequencyBackupScheduledEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
