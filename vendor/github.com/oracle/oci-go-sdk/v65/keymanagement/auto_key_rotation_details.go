// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoKeyRotationDetails The details of auto rotation schedule for the Key being create updated or imported.
type AutoKeyRotationDetails struct {

	// The interval of auto key rotation. For auto key rotation the interval should between 60 day and 365 days (1 year). Note: User must specify this parameter when creating a new schedule.
	RotationIntervalInDays *int `mandatory:"false" json:"rotationIntervalInDays"`

	// A property indicating  scheduled start date expressed as date YYYY-MM-DD String. Example: `2023-04-04T00:00:00Z. The time has no significance when scheduling an auto key rotation as this can be done anytime approximately the scheduled day, KMS ignores the time and replaces it with 00:00, for example 2023-04-04T15:14:13Z will be used as 2023-04-04T00:00:00Z . Note : Today’s date will be used if not specified by customer.
	TimeOfScheduleStart *common.SDKTime `mandatory:"false" json:"timeOfScheduleStart"`

	// A property indicating Next estimated scheduled Time, as per the interval, expressed as date YYYY-MM-DD String. Example: `2023-04-04T00:00:00Z`. The time has no significance when scheduling an auto key rotation as this can be done anytime approximately the scheduled day, KMS ignores the time and replaces it with 00:00, for example 2023-04-04T15:14:13Z will be used as 2023-04-04T00:00:00Z.
	TimeOfNextRotation *common.SDKTime `mandatory:"false" json:"timeOfNextRotation"`

	// A property indicating Last rotation Date. Example: `2023-04-04T00:00:00Z`.
	TimeOfLastRotation *common.SDKTime `mandatory:"false" json:"timeOfLastRotation"`

	// The status of last execution of auto key rotation.
	LastRotationStatus AutoKeyRotationDetailsLastRotationStatusEnum `mandatory:"false" json:"lastRotationStatus,omitempty"`

	// The last execution status message of auto key rotation.
	LastRotationMessage *string `mandatory:"false" json:"lastRotationMessage"`
}

func (m AutoKeyRotationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoKeyRotationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutoKeyRotationDetailsLastRotationStatusEnum(string(m.LastRotationStatus)); !ok && m.LastRotationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastRotationStatus: %s. Supported values are: %s.", m.LastRotationStatus, strings.Join(GetAutoKeyRotationDetailsLastRotationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoKeyRotationDetailsLastRotationStatusEnum Enum with underlying type: string
type AutoKeyRotationDetailsLastRotationStatusEnum string

// Set of constants representing the allowable values for AutoKeyRotationDetailsLastRotationStatusEnum
const (
	AutoKeyRotationDetailsLastRotationStatusSuccess    AutoKeyRotationDetailsLastRotationStatusEnum = "SUCCESS"
	AutoKeyRotationDetailsLastRotationStatusFailed     AutoKeyRotationDetailsLastRotationStatusEnum = "FAILED"
	AutoKeyRotationDetailsLastRotationStatusInProgress AutoKeyRotationDetailsLastRotationStatusEnum = "IN_PROGRESS"
)

var mappingAutoKeyRotationDetailsLastRotationStatusEnum = map[string]AutoKeyRotationDetailsLastRotationStatusEnum{
	"SUCCESS":     AutoKeyRotationDetailsLastRotationStatusSuccess,
	"FAILED":      AutoKeyRotationDetailsLastRotationStatusFailed,
	"IN_PROGRESS": AutoKeyRotationDetailsLastRotationStatusInProgress,
}

var mappingAutoKeyRotationDetailsLastRotationStatusEnumLowerCase = map[string]AutoKeyRotationDetailsLastRotationStatusEnum{
	"success":     AutoKeyRotationDetailsLastRotationStatusSuccess,
	"failed":      AutoKeyRotationDetailsLastRotationStatusFailed,
	"in_progress": AutoKeyRotationDetailsLastRotationStatusInProgress,
}

// GetAutoKeyRotationDetailsLastRotationStatusEnumValues Enumerates the set of values for AutoKeyRotationDetailsLastRotationStatusEnum
func GetAutoKeyRotationDetailsLastRotationStatusEnumValues() []AutoKeyRotationDetailsLastRotationStatusEnum {
	values := make([]AutoKeyRotationDetailsLastRotationStatusEnum, 0)
	for _, v := range mappingAutoKeyRotationDetailsLastRotationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoKeyRotationDetailsLastRotationStatusEnumStringValues Enumerates the set of values in String for AutoKeyRotationDetailsLastRotationStatusEnum
func GetAutoKeyRotationDetailsLastRotationStatusEnumStringValues() []string {
	return []string{
		"SUCCESS",
		"FAILED",
		"IN_PROGRESS",
	}
}

// GetMappingAutoKeyRotationDetailsLastRotationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoKeyRotationDetailsLastRotationStatusEnum(val string) (AutoKeyRotationDetailsLastRotationStatusEnum, bool) {
	enum, ok := mappingAutoKeyRotationDetailsLastRotationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
