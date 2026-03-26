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

// EnableAuditMgmtLogging Audit Log resource for tracking status and details of audit logging operations on HSM clusters.
type EnableAuditMgmtLogging struct {

	// Workflow request identifier.
	Id *string `mandatory:"true" json:"id"`

	// OCID of the HSM Cluster.
	HsmClusterId *string `mandatory:"true" json:"hsmClusterId"`

	// Name of the audit log's associated customer bucket.
	CustomerBucketName *string `mandatory:"true" json:"customerBucketName"`

	// Object Storage namespace of the bucket.
	Namespace *string `mandatory:"true" json:"namespace"`

	// Compartment OCID for the HSM cluster (bucket compartment).
	ClusterCompartmentId *string `mandatory:"true" json:"clusterCompartmentId"`

	// Status of the audit logging.
	AuditLoggingStatus EnableAuditMgmtLoggingAuditLoggingStatusEnum `mandatory:"true" json:"auditLoggingStatus"`
}

func (m EnableAuditMgmtLogging) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EnableAuditMgmtLogging) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEnableAuditMgmtLoggingAuditLoggingStatusEnum(string(m.AuditLoggingStatus)); !ok && m.AuditLoggingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditLoggingStatus: %s. Supported values are: %s.", m.AuditLoggingStatus, strings.Join(GetEnableAuditMgmtLoggingAuditLoggingStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EnableAuditMgmtLoggingAuditLoggingStatusEnum Enum with underlying type: string
type EnableAuditMgmtLoggingAuditLoggingStatusEnum string

// Set of constants representing the allowable values for EnableAuditMgmtLoggingAuditLoggingStatusEnum
const (
	EnableAuditMgmtLoggingAuditLoggingStatusCreating            EnableAuditMgmtLoggingAuditLoggingStatusEnum = "CREATING"
	EnableAuditMgmtLoggingAuditLoggingStatusWaitingForCustomer  EnableAuditMgmtLoggingAuditLoggingStatusEnum = "WAITING_FOR_CUSTOMER"
	EnableAuditMgmtLoggingAuditLoggingStatusVerifying           EnableAuditMgmtLoggingAuditLoggingStatusEnum = "VERIFYING"
	EnableAuditMgmtLoggingAuditLoggingStatusDisablingInProgress EnableAuditMgmtLoggingAuditLoggingStatusEnum = "DISABLING_IN_PROGRESS"
	EnableAuditMgmtLoggingAuditLoggingStatusDisabledValidated   EnableAuditMgmtLoggingAuditLoggingStatusEnum = "DISABLED_VALIDATED"
	EnableAuditMgmtLoggingAuditLoggingStatusEnabled             EnableAuditMgmtLoggingAuditLoggingStatusEnum = "ENABLED"
	EnableAuditMgmtLoggingAuditLoggingStatusDisabled            EnableAuditMgmtLoggingAuditLoggingStatusEnum = "DISABLED"
)

var mappingEnableAuditMgmtLoggingAuditLoggingStatusEnum = map[string]EnableAuditMgmtLoggingAuditLoggingStatusEnum{
	"CREATING":              EnableAuditMgmtLoggingAuditLoggingStatusCreating,
	"WAITING_FOR_CUSTOMER":  EnableAuditMgmtLoggingAuditLoggingStatusWaitingForCustomer,
	"VERIFYING":             EnableAuditMgmtLoggingAuditLoggingStatusVerifying,
	"DISABLING_IN_PROGRESS": EnableAuditMgmtLoggingAuditLoggingStatusDisablingInProgress,
	"DISABLED_VALIDATED":    EnableAuditMgmtLoggingAuditLoggingStatusDisabledValidated,
	"ENABLED":               EnableAuditMgmtLoggingAuditLoggingStatusEnabled,
	"DISABLED":              EnableAuditMgmtLoggingAuditLoggingStatusDisabled,
}

var mappingEnableAuditMgmtLoggingAuditLoggingStatusEnumLowerCase = map[string]EnableAuditMgmtLoggingAuditLoggingStatusEnum{
	"creating":              EnableAuditMgmtLoggingAuditLoggingStatusCreating,
	"waiting_for_customer":  EnableAuditMgmtLoggingAuditLoggingStatusWaitingForCustomer,
	"verifying":             EnableAuditMgmtLoggingAuditLoggingStatusVerifying,
	"disabling_in_progress": EnableAuditMgmtLoggingAuditLoggingStatusDisablingInProgress,
	"disabled_validated":    EnableAuditMgmtLoggingAuditLoggingStatusDisabledValidated,
	"enabled":               EnableAuditMgmtLoggingAuditLoggingStatusEnabled,
	"disabled":              EnableAuditMgmtLoggingAuditLoggingStatusDisabled,
}

// GetEnableAuditMgmtLoggingAuditLoggingStatusEnumValues Enumerates the set of values for EnableAuditMgmtLoggingAuditLoggingStatusEnum
func GetEnableAuditMgmtLoggingAuditLoggingStatusEnumValues() []EnableAuditMgmtLoggingAuditLoggingStatusEnum {
	values := make([]EnableAuditMgmtLoggingAuditLoggingStatusEnum, 0)
	for _, v := range mappingEnableAuditMgmtLoggingAuditLoggingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetEnableAuditMgmtLoggingAuditLoggingStatusEnumStringValues Enumerates the set of values in String for EnableAuditMgmtLoggingAuditLoggingStatusEnum
func GetEnableAuditMgmtLoggingAuditLoggingStatusEnumStringValues() []string {
	return []string{
		"CREATING",
		"WAITING_FOR_CUSTOMER",
		"VERIFYING",
		"DISABLING_IN_PROGRESS",
		"DISABLED_VALIDATED",
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingEnableAuditMgmtLoggingAuditLoggingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEnableAuditMgmtLoggingAuditLoggingStatusEnum(val string) (EnableAuditMgmtLoggingAuditLoggingStatusEnum, bool) {
	enum, ok := mappingEnableAuditMgmtLoggingAuditLoggingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
