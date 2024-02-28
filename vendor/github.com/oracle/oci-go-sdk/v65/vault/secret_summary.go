// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Secret Management API
//
// Use the Secret Management API to manage secrets and secret versions. For more information, see Managing Secrets (https://docs.cloud.oracle.com/Content/KeyManagement/Tasks/managingsecrets.htm).
//

package vault

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID of the master encryption key that is used to encrypt the secret. You must specify a symmetric key to encrypt the secret during import to the vault. You cannot encrypt secrets with asymmetric keys. Furthermore, the key must exist in the vault that you specify.
	KeyId *string `mandatory:"false" json:"keyId"`

	// Additional information about the secret's current lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	RotationConfig *RotationConfig `mandatory:"false" json:"rotationConfig"`

	// Additional information about the status of the secret rotation
	RotationStatus SecretRotationStatusEnum `mandatory:"false" json:"rotationStatus,omitempty"`

	// A property indicating when the secret was last rotated successfully, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	LastRotationTime *common.SDKTime `mandatory:"false" json:"lastRotationTime"`

	// A property indicating when the secret is scheduled to be rotated, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	NextRotationTime *common.SDKTime `mandatory:"false" json:"nextRotationTime"`

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

func (m SecretSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecretSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecretSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSecretSummaryLifecycleStateEnumStringValues(), ",")))
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
func (m *SecretSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DefinedTags                map[string]map[string]interface{} `json:"definedTags"`
		Description                *string                           `json:"description"`
		FreeformTags               map[string]string                 `json:"freeformTags"`
		SystemTags                 map[string]map[string]interface{} `json:"systemTags"`
		KeyId                      *string                           `json:"keyId"`
		LifecycleDetails           *string                           `json:"lifecycleDetails"`
		RotationConfig             *RotationConfig                   `json:"rotationConfig"`
		RotationStatus             SecretRotationStatusEnum          `json:"rotationStatus"`
		LastRotationTime           *common.SDKTime                   `json:"lastRotationTime"`
		NextRotationTime           *common.SDKTime                   `json:"nextRotationTime"`
		TimeOfCurrentVersionExpiry *common.SDKTime                   `json:"timeOfCurrentVersionExpiry"`
		TimeOfDeletion             *common.SDKTime                   `json:"timeOfDeletion"`
		SecretGenerationContext    secretgenerationcontext           `json:"secretGenerationContext"`
		IsAutoGenerationEnabled    *bool                             `json:"isAutoGenerationEnabled"`
		CompartmentId              *string                           `json:"compartmentId"`
		Id                         *string                           `json:"id"`
		LifecycleState             SecretSummaryLifecycleStateEnum   `json:"lifecycleState"`
		SecretName                 *string                           `json:"secretName"`
		TimeCreated                *common.SDKTime                   `json:"timeCreated"`
		VaultId                    *string                           `json:"vaultId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DefinedTags = model.DefinedTags

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.SystemTags = model.SystemTags

	m.KeyId = model.KeyId

	m.LifecycleDetails = model.LifecycleDetails

	m.RotationConfig = model.RotationConfig

	m.RotationStatus = model.RotationStatus

	m.LastRotationTime = model.LastRotationTime

	m.NextRotationTime = model.NextRotationTime

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

var mappingSecretSummaryLifecycleStateEnum = map[string]SecretSummaryLifecycleStateEnum{
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

var mappingSecretSummaryLifecycleStateEnumLowerCase = map[string]SecretSummaryLifecycleStateEnum{
	"creating":            SecretSummaryLifecycleStateCreating,
	"active":              SecretSummaryLifecycleStateActive,
	"updating":            SecretSummaryLifecycleStateUpdating,
	"deleting":            SecretSummaryLifecycleStateDeleting,
	"deleted":             SecretSummaryLifecycleStateDeleted,
	"scheduling_deletion": SecretSummaryLifecycleStateSchedulingDeletion,
	"pending_deletion":    SecretSummaryLifecycleStatePendingDeletion,
	"cancelling_deletion": SecretSummaryLifecycleStateCancellingDeletion,
	"failed":              SecretSummaryLifecycleStateFailed,
}

// GetSecretSummaryLifecycleStateEnumValues Enumerates the set of values for SecretSummaryLifecycleStateEnum
func GetSecretSummaryLifecycleStateEnumValues() []SecretSummaryLifecycleStateEnum {
	values := make([]SecretSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSecretSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SecretSummaryLifecycleStateEnum
func GetSecretSummaryLifecycleStateEnumStringValues() []string {
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

// GetMappingSecretSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretSummaryLifecycleStateEnum(val string) (SecretSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSecretSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
