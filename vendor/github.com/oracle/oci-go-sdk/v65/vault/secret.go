// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.oracle.com/iaas/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Secret The details of the secret. Secret details do not contain the contents of the secret itself.
type Secret struct {

	// The OCID of the compartment where you want to create the secret.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the secret.
	Id *string `mandatory:"true" json:"id"`

	// The current lifecycle state of the secret.
	LifecycleState SecretLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The user-friendly name of the secret. Avoid entering confidential information.
	SecretName *string `mandatory:"true" json:"secretName"`

	// A property indicating when the secret was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the vault where the secret exists.
	VaultId *string `mandatory:"true" json:"vaultId"`

	// The version number of the secret version that's currently in use.
	CurrentVersionNumber *int64 `mandatory:"false" json:"currentVersionNumber"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A brief description of the secret. Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The OCID of the master encryption key that is used to encrypt the secret. You must specify a symmetric key to encrypt the secret during import to the vault. You cannot encrypt secrets with asymmetric keys. Furthermore, the key must exist in the vault that you specify.
	KeyId *string `mandatory:"false" json:"keyId"`

	// Additional information about the current lifecycle state of the secret.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Additional metadata that you can use to provide context about how to use the secret or during rotation or
	// other administrative tasks. For example, for a secret that you use to connect to a database, the additional
	// metadata might specify the connection endpoint and the connection string. Provide additional metadata as key-value pairs.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`

	RotationConfig *RotationConfig `mandatory:"false" json:"rotationConfig"`

	// Additional information about the status of the secret rotation
	RotationStatus SecretRotationStatusEnum `mandatory:"false" json:"rotationStatus,omitempty"`

	// A property indicating when the secret was last rotated successfully, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	LastRotationTime *common.SDKTime `mandatory:"false" json:"lastRotationTime"`

	// A property indicating when the secret is scheduled to be rotated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	NextRotationTime *common.SDKTime `mandatory:"false" json:"nextRotationTime"`

	// A list of rules that control how the secret is used and managed.
	SecretRules []SecretRule `mandatory:"false" json:"secretRules"`

	// An optional property indicating when the current secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfCurrentVersionExpiry *common.SDKTime `mandatory:"false" json:"timeOfCurrentVersionExpiry"`

	// An optional property indicating when to delete the secret, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	SecretGenerationContext SecretGenerationContext `mandatory:"false" json:"secretGenerationContext"`

	// The value of this flag determines whether or not secret content will be generated automatically.
	IsAutoGenerationEnabled *bool `mandatory:"false" json:"isAutoGenerationEnabled"`
}

func (m Secret) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Secret) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecretLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecretLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSecretRotationStatusEnum(string(m.RotationStatus)); !ok && m.RotationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RotationStatus: %s. Supported values are: %s.", m.RotationStatus, strings.Join(GetSecretRotationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *Secret) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CurrentVersionNumber       *int64                            `json:"currentVersionNumber"`
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		Description                *string                           `json:"description"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		KeyId                      *string                           `json:"keyId"`
		LifecycleDetails           *string                           `json:"lifecycleDetails"`
		Metadata                   map[string]interface{}            `json:"metadata"`
		RotationConfig             *RotationConfig                   `json:"rotationConfig"`
		RotationStatus             SecretRotationStatusEnum          `json:"rotationStatus"`
		LastRotationTime           *common.SDKTime                   `json:"lastRotationTime"`
		NextRotationTime           *common.SDKTime                   `json:"nextRotationTime"`
		SecretRules                []secretrule                      `json:"secretRules"`
		TimeOfCurrentVersionExpiry *common.SDKTime                   `json:"timeOfCurrentVersionExpiry"`
		TimeOfDeletion             *common.SDKTime                   `json:"timeOfDeletion"`
		SecretGenerationContext    secretgenerationcontext           `json:"secretGenerationContext"`
		IsAutoGenerationEnabled    *bool                             `json:"isAutoGenerationEnabled"`
		CompartmentId              *string                           `json:"compartmentId"`
		Id                         *string                           `json:"id"`
		LifecycleState             SecretLifecycleStateEnum          `json:"lifecycleState"`
		SecretName                 *string                           `json:"secretName"`
		TimeCreated                *common.SDKTime                   `json:"timeCreated"`
		VaultId                    *string                           `json:"vaultId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CurrentVersionNumber = model.CurrentVersionNumber

	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.KeyId = model.KeyId

	m.LifecycleDetails = model.LifecycleDetails

	m.Metadata = model.Metadata

	m.RotationConfig = model.RotationConfig

	m.RotationStatus = model.RotationStatus

	m.LastRotationTime = model.LastRotationTime

	m.NextRotationTime = model.NextRotationTime

	m.SecretRules = make([]SecretRule, len(model.SecretRules))
	for i, n := range model.SecretRules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.SecretRules[i] = nn.(SecretRule)
		} else {
			m.SecretRules[i] = nil
		}
	}
	m.TimeOfCurrentVersionExpiry = model.TimeOfCurrentVersionExpiry

	m.TimeOfDeletion = model.TimeOfDeletion

	nn, e = model.SecretGenerationContext.UnmarshalPolymorphicJSON(model.SecretGenerationContext.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretGenerationContext = nn.(SecretGenerationContext)
	} else {
		m.SecretGenerationContext = nil
	}

	m.IsAutoGenerationEnabled = model.IsAutoGenerationEnabled

	m.CompartmentId = model.CompartmentId

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	m.SecretName = model.SecretName

	m.TimeCreated = model.TimeCreated

	m.VaultId = model.VaultId

	return
}

// SecretLifecycleStateEnum Enum with underlying type: string
type SecretLifecycleStateEnum string

// Set of constants representing the allowable values for SecretLifecycleStateEnum
const (
	SecretLifecycleStateCreating           SecretLifecycleStateEnum = "CREATING"
	SecretLifecycleStateActive             SecretLifecycleStateEnum = "ACTIVE"
	SecretLifecycleStateUpdating           SecretLifecycleStateEnum = "UPDATING"
	SecretLifecycleStateDeleting           SecretLifecycleStateEnum = "DELETING"
	SecretLifecycleStateDeleted            SecretLifecycleStateEnum = "DELETED"
	SecretLifecycleStateSchedulingDeletion SecretLifecycleStateEnum = "SCHEDULING_DELETION"
	SecretLifecycleStatePendingDeletion    SecretLifecycleStateEnum = "PENDING_DELETION"
	SecretLifecycleStateCancellingDeletion SecretLifecycleStateEnum = "CANCELLING_DELETION"
	SecretLifecycleStateFailed             SecretLifecycleStateEnum = "FAILED"
)

var mappingSecretLifecycleStateEnum = map[string]SecretLifecycleStateEnum{
	"CREATING":            SecretLifecycleStateCreating,
	"ACTIVE":              SecretLifecycleStateActive,
	"UPDATING":            SecretLifecycleStateUpdating,
	"DELETING":            SecretLifecycleStateDeleting,
	"DELETED":             SecretLifecycleStateDeleted,
	"SCHEDULING_DELETION": SecretLifecycleStateSchedulingDeletion,
	"PENDING_DELETION":    SecretLifecycleStatePendingDeletion,
	"CANCELLING_DELETION": SecretLifecycleStateCancellingDeletion,
	"FAILED":              SecretLifecycleStateFailed,
}

var mappingSecretLifecycleStateEnumLowerCase = map[string]SecretLifecycleStateEnum{
	"creating":            SecretLifecycleStateCreating,
	"active":              SecretLifecycleStateActive,
	"updating":            SecretLifecycleStateUpdating,
	"deleting":            SecretLifecycleStateDeleting,
	"deleted":             SecretLifecycleStateDeleted,
	"scheduling_deletion": SecretLifecycleStateSchedulingDeletion,
	"pending_deletion":    SecretLifecycleStatePendingDeletion,
	"cancelling_deletion": SecretLifecycleStateCancellingDeletion,
	"failed":              SecretLifecycleStateFailed,
}

// GetSecretLifecycleStateEnumValues Enumerates the set of values for SecretLifecycleStateEnum
func GetSecretLifecycleStateEnumValues() []SecretLifecycleStateEnum {
	values := make([]SecretLifecycleStateEnum, 0)
	for _, v := range mappingSecretLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretLifecycleStateEnumStringValues Enumerates the set of values in String for SecretLifecycleStateEnum
func GetSecretLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"SCHEDULING_DELETION",
		"PENDING_DELETION",
		"CANCELLING_DELETION",
		"FAILED",
	}
}

// GetMappingSecretLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretLifecycleStateEnum(val string) (SecretLifecycleStateEnum, bool) {
	enum, ok := mappingSecretLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecretRotationStatusEnum Enum with underlying type: string
type SecretRotationStatusEnum string

// Set of constants representing the allowable values for SecretRotationStatusEnum
const (
	SecretRotationStatusInProgress SecretRotationStatusEnum = "IN_PROGRESS"
	SecretRotationStatusScheduled  SecretRotationStatusEnum = "SCHEDULED"
	SecretRotationStatusNotEnabled SecretRotationStatusEnum = "NOT_ENABLED"
	SecretRotationStatusCancelling SecretRotationStatusEnum = "CANCELLING"
)

var mappingSecretRotationStatusEnum = map[string]SecretRotationStatusEnum{
	"IN_PROGRESS": SecretRotationStatusInProgress,
	"SCHEDULED":   SecretRotationStatusScheduled,
	"NOT_ENABLED": SecretRotationStatusNotEnabled,
	"CANCELLING":  SecretRotationStatusCancelling,
}

var mappingSecretRotationStatusEnumLowerCase = map[string]SecretRotationStatusEnum{
	"in_progress": SecretRotationStatusInProgress,
	"scheduled":   SecretRotationStatusScheduled,
	"not_enabled": SecretRotationStatusNotEnabled,
	"cancelling":  SecretRotationStatusCancelling,
}

// GetSecretRotationStatusEnumValues Enumerates the set of values for SecretRotationStatusEnum
func GetSecretRotationStatusEnumValues() []SecretRotationStatusEnum {
	values := make([]SecretRotationStatusEnum, 0)
	for _, v := range mappingSecretRotationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretRotationStatusEnumStringValues Enumerates the set of values in String for SecretRotationStatusEnum
func GetSecretRotationStatusEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"SCHEDULED",
		"NOT_ENABLED",
		"CANCELLING",
	}
}

// GetMappingSecretRotationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretRotationStatusEnum(val string) (SecretRotationStatusEnum, bool) {
	enum, ok := mappingSecretRotationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
