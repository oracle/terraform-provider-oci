// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets Management API
//
// API for managing secrets.
//

package vault

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SecretSummary The details of the secret, excluding the contents of the secret.
type SecretSummary struct {

	// The OCID of the compartment that contains the secret.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the secret.
	Id *string `mandatory:"true" json:"id"`

	// The current lifecycle state of the secret.
	LifecycleState SecretSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the secret.
	SecretName *string `mandatory:"true" json:"secretName"`

	// A property indicating when the secret was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the Vault in which the secret exists
	VaultId *string `mandatory:"true" json:"vaultId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A brief description of the secret.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the master encryption key that is used to encrypt the secret.
	KeyId *string `mandatory:"false" json:"keyId"`

	// Additional information about the secret's current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// An optional property indicating when the current secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfCurrentVersionExpiry *common.SDKTime `mandatory:"false" json:"timeOfCurrentVersionExpiry"`

	// An optional property indicating when to delete the secret, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`
}

func (m SecretSummary) String() string {
	return common.PointerString(m)
}

// SecretSummaryLifecycleStateEnum Enum with underlying type: string
type SecretSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SecretSummaryLifecycleStateEnum
const (
	SecretSummaryLifecycleStateCreating           SecretSummaryLifecycleStateEnum = "CREATING"
	SecretSummaryLifecycleStateActive             SecretSummaryLifecycleStateEnum = "ACTIVE"
	SecretSummaryLifecycleStateUpdating           SecretSummaryLifecycleStateEnum = "UPDATING"
	SecretSummaryLifecycleStateDeleting           SecretSummaryLifecycleStateEnum = "DELETING"
	SecretSummaryLifecycleStateDeleted            SecretSummaryLifecycleStateEnum = "DELETED"
	SecretSummaryLifecycleStateSchedulingDeletion SecretSummaryLifecycleStateEnum = "SCHEDULING_DELETION"
	SecretSummaryLifecycleStatePendingDeletion    SecretSummaryLifecycleStateEnum = "PENDING_DELETION"
	SecretSummaryLifecycleStateCancellingDeletion SecretSummaryLifecycleStateEnum = "CANCELLING_DELETION"
	SecretSummaryLifecycleStateFailed             SecretSummaryLifecycleStateEnum = "FAILED"
)

var mappingSecretSummaryLifecycleState = map[string]SecretSummaryLifecycleStateEnum{
	"CREATING":            SecretSummaryLifecycleStateCreating,
	"ACTIVE":              SecretSummaryLifecycleStateActive,
	"UPDATING":            SecretSummaryLifecycleStateUpdating,
	"DELETING":            SecretSummaryLifecycleStateDeleting,
	"DELETED":             SecretSummaryLifecycleStateDeleted,
	"SCHEDULING_DELETION": SecretSummaryLifecycleStateSchedulingDeletion,
	"PENDING_DELETION":    SecretSummaryLifecycleStatePendingDeletion,
	"CANCELLING_DELETION": SecretSummaryLifecycleStateCancellingDeletion,
	"FAILED":              SecretSummaryLifecycleStateFailed,
}

// GetSecretSummaryLifecycleStateEnumValues Enumerates the set of values for SecretSummaryLifecycleStateEnum
func GetSecretSummaryLifecycleStateEnumValues() []SecretSummaryLifecycleStateEnum {
	values := make([]SecretSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSecretSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
