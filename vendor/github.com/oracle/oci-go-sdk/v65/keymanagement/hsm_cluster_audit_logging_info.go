// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Key Management API
//
// Use the Key Management API to manage vaults and keys. For more information, see Managing Vaults (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingvaults.htm) and Managing Keys (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingkeys.htm).
//

package keymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// HsmClusterAuditLoggingInfo Status of management audit logging for the cluster.
type HsmClusterAuditLoggingInfo struct {

	// Whether audit logging is enabled.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// Audit logs bucket name.
	BucketName *string `mandatory:"true" json:"bucketName"`

	// Bucket namespace.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Compartment of the bucket.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of dynamic group used for audit log upload.
	DynamicGroupName *string `mandatory:"true" json:"dynamicGroupName"`

	// The current lifecycle state of the audit logs.
	AuditLogLifecycleState HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum `mandatory:"true" json:"auditLogLifecycleState"`
}

func (m HsmClusterAuditLoggingInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HsmClusterAuditLoggingInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum(string(m.AuditLogLifecycleState)); !ok && m.AuditLogLifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditLogLifecycleState: %s. Supported values are: %s.", m.AuditLogLifecycleState, strings.Join(GetHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum Enum with underlying type: string
type HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum string

// Set of constants representing the allowable values for HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum
const (
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateCreating            HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "CREATING"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateWaitingForCustomer  HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "WAITING_FOR_CUSTOMER"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateVerifying           HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "VERIFYING"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisablingInProgress HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "DISABLING_IN_PROGRESS"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisablingBroken     HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "DISABLING_BROKEN"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisabledValidated   HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "DISABLED_VALIDATED"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateFailed              HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "FAILED"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnabled             HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "ENABLED"
	HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisabled            HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = "DISABLED"
)

var mappingHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum = map[string]HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum{
	"CREATING":              HsmClusterAuditLoggingInfoAuditLogLifecycleStateCreating,
	"WAITING_FOR_CUSTOMER":  HsmClusterAuditLoggingInfoAuditLogLifecycleStateWaitingForCustomer,
	"VERIFYING":             HsmClusterAuditLoggingInfoAuditLogLifecycleStateVerifying,
	"DISABLING_IN_PROGRESS": HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisablingInProgress,
	"DISABLING_BROKEN":      HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisablingBroken,
	"DISABLED_VALIDATED":    HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisabledValidated,
	"FAILED":                HsmClusterAuditLoggingInfoAuditLogLifecycleStateFailed,
	"ENABLED":               HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnabled,
	"DISABLED":              HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisabled,
}

var mappingHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnumLowerCase = map[string]HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum{
	"creating":              HsmClusterAuditLoggingInfoAuditLogLifecycleStateCreating,
	"waiting_for_customer":  HsmClusterAuditLoggingInfoAuditLogLifecycleStateWaitingForCustomer,
	"verifying":             HsmClusterAuditLoggingInfoAuditLogLifecycleStateVerifying,
	"disabling_in_progress": HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisablingInProgress,
	"disabling_broken":      HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisablingBroken,
	"disabled_validated":    HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisabledValidated,
	"failed":                HsmClusterAuditLoggingInfoAuditLogLifecycleStateFailed,
	"enabled":               HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnabled,
	"disabled":              HsmClusterAuditLoggingInfoAuditLogLifecycleStateDisabled,
}

// GetHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnumValues Enumerates the set of values for HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum
func GetHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnumValues() []HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum {
	values := make([]HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum, 0)
	for _, v := range mappingHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnumStringValues Enumerates the set of values in String for HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum
func GetHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"WAITING_FOR_CUSTOMER",
		"VERIFYING",
		"DISABLING_IN_PROGRESS",
		"DISABLING_BROKEN",
		"DISABLED_VALIDATED",
		"FAILED",
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum(val string) (HsmClusterAuditLoggingInfoAuditLogLifecycleStateEnum, bool) {
	enum, ok := mappingHsmClusterAuditLoggingInfoAuditLogLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
