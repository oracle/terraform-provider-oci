// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Digital Assistant Service Instance API
//
// API to create and maintain Oracle Digital Assistant service instances.
//

package oda

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OdaInstanceSummary Summary of the Digital Assistant instance.
type OdaInstanceSummary struct {

	// Unique identifier of the Digital Assistant instance.
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment that the instance belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the instance.
	LifecycleState OdaInstanceSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// User-defined name for the Digital Assistant instance. You can change this value.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Digital Assistant instance.
	Description *string `mandatory:"false" json:"description"`

	// Shape or size of the instance.
	ShapeName OdaInstanceSummaryShapeNameEnum `mandatory:"false" json:"shapeName,omitempty"`

	// When the Digital Assistant instance was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the Digital Assistant instance was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current sub-state of the Digital Assistant instance.
	LifecycleSubState OdaInstanceSummaryLifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// A message describing the current state in more detail. For example, actionable
	// information about an instance that's in the `FAILED` state.
	StateMessage *string `mandatory:"false" json:"stateMessage"`

	// Simple key-value pair that is applied without any predefined name, type, or scope.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Should this Digital Assistant instance use role-based authorization via an identity domain (true) or use the default policy-based authorization via IAM policies (false)
	IsRoleBasedAccess *bool `mandatory:"false" json:"isRoleBasedAccess"`

	// If isRoleBasedAccess is set to true, this property specifies the identity domain that is to be used to implement this type of authorzation. Digital Assistant will create an Identity Application instance and Application Roles within this identity domain. The caller may then perform and user roll mappings they like to grant access to users within the identity domain.
	IdentityDomain *string `mandatory:"false" json:"identityDomain"`

	// A list of package names imported into this instance (if any).
	ImportedPackageNames []string `mandatory:"false" json:"importedPackageNames"`

	// A list of attachment types for this instance (if any).
	AttachmentTypes []string `mandatory:"false" json:"attachmentTypes"`
}

func (m OdaInstanceSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OdaInstanceSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaInstanceSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOdaInstanceSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOdaInstanceSummaryShapeNameEnum(string(m.ShapeName)); !ok && m.ShapeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeName: %s. Supported values are: %s.", m.ShapeName, strings.Join(GetOdaInstanceSummaryShapeNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOdaInstanceSummaryLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetOdaInstanceSummaryLifecycleSubStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdaInstanceSummaryShapeNameEnum Enum with underlying type: string
type OdaInstanceSummaryShapeNameEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryShapeNameEnum
const (
	OdaInstanceSummaryShapeNameDevelopment OdaInstanceSummaryShapeNameEnum = "DEVELOPMENT"
	OdaInstanceSummaryShapeNameProduction  OdaInstanceSummaryShapeNameEnum = "PRODUCTION"
)

var mappingOdaInstanceSummaryShapeNameEnum = map[string]OdaInstanceSummaryShapeNameEnum{
	"DEVELOPMENT": OdaInstanceSummaryShapeNameDevelopment,
	"PRODUCTION":  OdaInstanceSummaryShapeNameProduction,
}

var mappingOdaInstanceSummaryShapeNameEnumLowerCase = map[string]OdaInstanceSummaryShapeNameEnum{
	"development": OdaInstanceSummaryShapeNameDevelopment,
	"production":  OdaInstanceSummaryShapeNameProduction,
}

// GetOdaInstanceSummaryShapeNameEnumValues Enumerates the set of values for OdaInstanceSummaryShapeNameEnum
func GetOdaInstanceSummaryShapeNameEnumValues() []OdaInstanceSummaryShapeNameEnum {
	values := make([]OdaInstanceSummaryShapeNameEnum, 0)
	for _, v := range mappingOdaInstanceSummaryShapeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaInstanceSummaryShapeNameEnumStringValues Enumerates the set of values in String for OdaInstanceSummaryShapeNameEnum
func GetOdaInstanceSummaryShapeNameEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingOdaInstanceSummaryShapeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaInstanceSummaryShapeNameEnum(val string) (OdaInstanceSummaryShapeNameEnum, bool) {
	enum, ok := mappingOdaInstanceSummaryShapeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OdaInstanceSummaryLifecycleStateEnum Enum with underlying type: string
type OdaInstanceSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryLifecycleStateEnum
const (
	OdaInstanceSummaryLifecycleStateCreating OdaInstanceSummaryLifecycleStateEnum = "CREATING"
	OdaInstanceSummaryLifecycleStateUpdating OdaInstanceSummaryLifecycleStateEnum = "UPDATING"
	OdaInstanceSummaryLifecycleStateActive   OdaInstanceSummaryLifecycleStateEnum = "ACTIVE"
	OdaInstanceSummaryLifecycleStateInactive OdaInstanceSummaryLifecycleStateEnum = "INACTIVE"
	OdaInstanceSummaryLifecycleStateDeleting OdaInstanceSummaryLifecycleStateEnum = "DELETING"
	OdaInstanceSummaryLifecycleStateDeleted  OdaInstanceSummaryLifecycleStateEnum = "DELETED"
	OdaInstanceSummaryLifecycleStateFailed   OdaInstanceSummaryLifecycleStateEnum = "FAILED"
)

var mappingOdaInstanceSummaryLifecycleStateEnum = map[string]OdaInstanceSummaryLifecycleStateEnum{
	"CREATING": OdaInstanceSummaryLifecycleStateCreating,
	"UPDATING": OdaInstanceSummaryLifecycleStateUpdating,
	"ACTIVE":   OdaInstanceSummaryLifecycleStateActive,
	"INACTIVE": OdaInstanceSummaryLifecycleStateInactive,
	"DELETING": OdaInstanceSummaryLifecycleStateDeleting,
	"DELETED":  OdaInstanceSummaryLifecycleStateDeleted,
	"FAILED":   OdaInstanceSummaryLifecycleStateFailed,
}

var mappingOdaInstanceSummaryLifecycleStateEnumLowerCase = map[string]OdaInstanceSummaryLifecycleStateEnum{
	"creating": OdaInstanceSummaryLifecycleStateCreating,
	"updating": OdaInstanceSummaryLifecycleStateUpdating,
	"active":   OdaInstanceSummaryLifecycleStateActive,
	"inactive": OdaInstanceSummaryLifecycleStateInactive,
	"deleting": OdaInstanceSummaryLifecycleStateDeleting,
	"deleted":  OdaInstanceSummaryLifecycleStateDeleted,
	"failed":   OdaInstanceSummaryLifecycleStateFailed,
}

// GetOdaInstanceSummaryLifecycleStateEnumValues Enumerates the set of values for OdaInstanceSummaryLifecycleStateEnum
func GetOdaInstanceSummaryLifecycleStateEnumValues() []OdaInstanceSummaryLifecycleStateEnum {
	values := make([]OdaInstanceSummaryLifecycleStateEnum, 0)
	for _, v := range mappingOdaInstanceSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaInstanceSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for OdaInstanceSummaryLifecycleStateEnum
func GetOdaInstanceSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"INACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOdaInstanceSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaInstanceSummaryLifecycleStateEnum(val string) (OdaInstanceSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingOdaInstanceSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OdaInstanceSummaryLifecycleSubStateEnum Enum with underlying type: string
type OdaInstanceSummaryLifecycleSubStateEnum string

// Set of constants representing the allowable values for OdaInstanceSummaryLifecycleSubStateEnum
const (
	OdaInstanceSummaryLifecycleSubStateCreating                          OdaInstanceSummaryLifecycleSubStateEnum = "CREATING"
	OdaInstanceSummaryLifecycleSubStateStarting                          OdaInstanceSummaryLifecycleSubStateEnum = "STARTING"
	OdaInstanceSummaryLifecycleSubStateStopping                          OdaInstanceSummaryLifecycleSubStateEnum = "STOPPING"
	OdaInstanceSummaryLifecycleSubStateChangingCompartment               OdaInstanceSummaryLifecycleSubStateEnum = "CHANGING_COMPARTMENT"
	OdaInstanceSummaryLifecycleSubStateActivatingCustomerEncryptionKey   OdaInstanceSummaryLifecycleSubStateEnum = "ACTIVATING_CUSTOMER_ENCRYPTION_KEY"
	OdaInstanceSummaryLifecycleSubStateUpdatingCustomerEncryptionKey     OdaInstanceSummaryLifecycleSubStateEnum = "UPDATING_CUSTOMER_ENCRYPTION_KEY"
	OdaInstanceSummaryLifecycleSubStateDeactivatingCustomerEncryptionKey OdaInstanceSummaryLifecycleSubStateEnum = "DEACTIVATING_CUSTOMER_ENCRYPTION_KEY"
	OdaInstanceSummaryLifecycleSubStateDeleting                          OdaInstanceSummaryLifecycleSubStateEnum = "DELETING"
	OdaInstanceSummaryLifecycleSubStateDeletePending                     OdaInstanceSummaryLifecycleSubStateEnum = "DELETE_PENDING"
	OdaInstanceSummaryLifecycleSubStateRecovering                        OdaInstanceSummaryLifecycleSubStateEnum = "RECOVERING"
	OdaInstanceSummaryLifecycleSubStateUpdating                          OdaInstanceSummaryLifecycleSubStateEnum = "UPDATING"
	OdaInstanceSummaryLifecycleSubStatePurging                           OdaInstanceSummaryLifecycleSubStateEnum = "PURGING"
	OdaInstanceSummaryLifecycleSubStateQueued                            OdaInstanceSummaryLifecycleSubStateEnum = "QUEUED"
)

var mappingOdaInstanceSummaryLifecycleSubStateEnum = map[string]OdaInstanceSummaryLifecycleSubStateEnum{
	"CREATING":                             OdaInstanceSummaryLifecycleSubStateCreating,
	"STARTING":                             OdaInstanceSummaryLifecycleSubStateStarting,
	"STOPPING":                             OdaInstanceSummaryLifecycleSubStateStopping,
	"CHANGING_COMPARTMENT":                 OdaInstanceSummaryLifecycleSubStateChangingCompartment,
	"ACTIVATING_CUSTOMER_ENCRYPTION_KEY":   OdaInstanceSummaryLifecycleSubStateActivatingCustomerEncryptionKey,
	"UPDATING_CUSTOMER_ENCRYPTION_KEY":     OdaInstanceSummaryLifecycleSubStateUpdatingCustomerEncryptionKey,
	"DEACTIVATING_CUSTOMER_ENCRYPTION_KEY": OdaInstanceSummaryLifecycleSubStateDeactivatingCustomerEncryptionKey,
	"DELETING":                             OdaInstanceSummaryLifecycleSubStateDeleting,
	"DELETE_PENDING":                       OdaInstanceSummaryLifecycleSubStateDeletePending,
	"RECOVERING":                           OdaInstanceSummaryLifecycleSubStateRecovering,
	"UPDATING":                             OdaInstanceSummaryLifecycleSubStateUpdating,
	"PURGING":                              OdaInstanceSummaryLifecycleSubStatePurging,
	"QUEUED":                               OdaInstanceSummaryLifecycleSubStateQueued,
}

var mappingOdaInstanceSummaryLifecycleSubStateEnumLowerCase = map[string]OdaInstanceSummaryLifecycleSubStateEnum{
	"creating":                             OdaInstanceSummaryLifecycleSubStateCreating,
	"starting":                             OdaInstanceSummaryLifecycleSubStateStarting,
	"stopping":                             OdaInstanceSummaryLifecycleSubStateStopping,
	"changing_compartment":                 OdaInstanceSummaryLifecycleSubStateChangingCompartment,
	"activating_customer_encryption_key":   OdaInstanceSummaryLifecycleSubStateActivatingCustomerEncryptionKey,
	"updating_customer_encryption_key":     OdaInstanceSummaryLifecycleSubStateUpdatingCustomerEncryptionKey,
	"deactivating_customer_encryption_key": OdaInstanceSummaryLifecycleSubStateDeactivatingCustomerEncryptionKey,
	"deleting":                             OdaInstanceSummaryLifecycleSubStateDeleting,
	"delete_pending":                       OdaInstanceSummaryLifecycleSubStateDeletePending,
	"recovering":                           OdaInstanceSummaryLifecycleSubStateRecovering,
	"updating":                             OdaInstanceSummaryLifecycleSubStateUpdating,
	"purging":                              OdaInstanceSummaryLifecycleSubStatePurging,
	"queued":                               OdaInstanceSummaryLifecycleSubStateQueued,
}

// GetOdaInstanceSummaryLifecycleSubStateEnumValues Enumerates the set of values for OdaInstanceSummaryLifecycleSubStateEnum
func GetOdaInstanceSummaryLifecycleSubStateEnumValues() []OdaInstanceSummaryLifecycleSubStateEnum {
	values := make([]OdaInstanceSummaryLifecycleSubStateEnum, 0)
	for _, v := range mappingOdaInstanceSummaryLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaInstanceSummaryLifecycleSubStateEnumStringValues Enumerates the set of values in String for OdaInstanceSummaryLifecycleSubStateEnum
func GetOdaInstanceSummaryLifecycleSubStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"STARTING",
		"STOPPING",
		"CHANGING_COMPARTMENT",
		"ACTIVATING_CUSTOMER_ENCRYPTION_KEY",
		"UPDATING_CUSTOMER_ENCRYPTION_KEY",
		"DEACTIVATING_CUSTOMER_ENCRYPTION_KEY",
		"DELETING",
		"DELETE_PENDING",
		"RECOVERING",
		"UPDATING",
		"PURGING",
		"QUEUED",
	}
}

// GetMappingOdaInstanceSummaryLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaInstanceSummaryLifecycleSubStateEnum(val string) (OdaInstanceSummaryLifecycleSubStateEnum, bool) {
	enum, ok := mappingOdaInstanceSummaryLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
