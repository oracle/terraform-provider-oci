// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KmsKeyInfo kmsKeyInfo
type KmsKeyInfo struct {

	// current BYOK keyId facp is using
	ActiveKeyId *string `mandatory:"false" json:"activeKeyId"`

	// current key version facp is using
	ActiveKeyVersion *string `mandatory:"false" json:"activeKeyVersion"`

	// scheduled keyId to be updated
	ScheduledKeyId *string `mandatory:"false" json:"scheduledKeyId"`

	// scheduled key version to be updated.
	ScheduledKeyVersion *string `mandatory:"false" json:"scheduledKeyVersion"`

	// current key lifeCycleState
	CurrentKeyLifecycleState *string `mandatory:"false" json:"currentKeyLifecycleState"`

	// scheduled key lifeCycle state to be updated.
	ScheduledLifecycleState *string `mandatory:"false" json:"scheduledLifecycleState"`

	// the scheduled key status
	ScheduledKeyStatus KmsKeyInfoScheduledKeyStatusEnum `mandatory:"false" json:"scheduledKeyStatus,omitempty"`
}

func (m KmsKeyInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KmsKeyInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingKmsKeyInfoScheduledKeyStatusEnum(string(m.ScheduledKeyStatus)); !ok && m.ScheduledKeyStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ScheduledKeyStatus: %s. Supported values are: %s.", m.ScheduledKeyStatus, strings.Join(GetKmsKeyInfoScheduledKeyStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// KmsKeyInfoScheduledKeyStatusEnum Enum with underlying type: string
type KmsKeyInfoScheduledKeyStatusEnum string

// Set of constants representing the allowable values for KmsKeyInfoScheduledKeyStatusEnum
const (
	KmsKeyInfoScheduledKeyStatusScheduling KmsKeyInfoScheduledKeyStatusEnum = "SCHEDULING"
	KmsKeyInfoScheduledKeyStatusUpdating   KmsKeyInfoScheduledKeyStatusEnum = "UPDATING"
	KmsKeyInfoScheduledKeyStatusFailed     KmsKeyInfoScheduledKeyStatusEnum = "FAILED"
	KmsKeyInfoScheduledKeyStatusNone       KmsKeyInfoScheduledKeyStatusEnum = "NONE"
)

var mappingKmsKeyInfoScheduledKeyStatusEnum = map[string]KmsKeyInfoScheduledKeyStatusEnum{
	"SCHEDULING": KmsKeyInfoScheduledKeyStatusScheduling,
	"UPDATING":   KmsKeyInfoScheduledKeyStatusUpdating,
	"FAILED":     KmsKeyInfoScheduledKeyStatusFailed,
	"NONE":       KmsKeyInfoScheduledKeyStatusNone,
}

var mappingKmsKeyInfoScheduledKeyStatusEnumLowerCase = map[string]KmsKeyInfoScheduledKeyStatusEnum{
	"scheduling": KmsKeyInfoScheduledKeyStatusScheduling,
	"updating":   KmsKeyInfoScheduledKeyStatusUpdating,
	"failed":     KmsKeyInfoScheduledKeyStatusFailed,
	"none":       KmsKeyInfoScheduledKeyStatusNone,
}

// GetKmsKeyInfoScheduledKeyStatusEnumValues Enumerates the set of values for KmsKeyInfoScheduledKeyStatusEnum
func GetKmsKeyInfoScheduledKeyStatusEnumValues() []KmsKeyInfoScheduledKeyStatusEnum {
	values := make([]KmsKeyInfoScheduledKeyStatusEnum, 0)
	for _, v := range mappingKmsKeyInfoScheduledKeyStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetKmsKeyInfoScheduledKeyStatusEnumStringValues Enumerates the set of values in String for KmsKeyInfoScheduledKeyStatusEnum
func GetKmsKeyInfoScheduledKeyStatusEnumStringValues() []string {
	return []string{
		"SCHEDULING",
		"UPDATING",
		"FAILED",
		"NONE",
	}
}

// GetMappingKmsKeyInfoScheduledKeyStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingKmsKeyInfoScheduledKeyStatusEnum(val string) (KmsKeyInfoScheduledKeyStatusEnum, bool) {
	enum, ok := mappingKmsKeyInfoScheduledKeyStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
