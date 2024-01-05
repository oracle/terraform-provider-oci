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

// OdaInstance Description of `OdaServiceInstance` object.
type OdaInstance struct {

	// Unique immutable identifier that was assigned when the instance was created.
	Id *string `mandatory:"true" json:"id"`

	// Identifier of the compartment that the instance belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Shape or size of the instance.
	ShapeName OdaInstanceShapeNameEnum `mandatory:"true" json:"shapeName"`

	// User-defined name for the Digital Assistant instance. Avoid entering confidential information.
	// You can change this value.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Description of the Digital Assistant instance.
	Description *string `mandatory:"false" json:"description"`

	// URL for the Digital Assistant web application that's associated with the instance.
	WebAppUrl *string `mandatory:"false" json:"webAppUrl"`

	// URL for the connector's endpoint.
	ConnectorUrl *string `mandatory:"false" json:"connectorUrl"`

	// When the Digital Assistant instance was created. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// When the Digital Assistance instance was last updated. A date-time string as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the Digital Assistant instance.
	LifecycleState OdaInstanceLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The current sub-state of the Digital Assistant instance.
	LifecycleSubState OdaInstanceLifecycleSubStateEnum `mandatory:"false" json:"lifecycleSubState,omitempty"`

	// A message that describes the current state in more detail.
	// For example, actionable information about an instance that's in the `FAILED` state.
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

	// If isRoleBasedAccess is set to true, this property specifies the GUID of the Identity Application instance Digital Assistant has created inside the user-specified identity domain. This identity application instance may be used to host user roll mappings to grant access to this Digital Assistant instance for users within the identity domain.
	IdentityAppGuid *string `mandatory:"false" json:"identityAppGuid"`

	// If isRoleBasedAccess is set to true, this property specifies the URL for the administration console used to manage the Identity Application instance Digital Assistant has created inside the user-specified identity domain.
	IdentityAppConsoleUrl *string `mandatory:"false" json:"identityAppConsoleUrl"`

	// A list of package names imported into this instance (if any). Use importedPackageIds field to get the details of the imported packages.
	ImportedPackageNames []string `mandatory:"false" json:"importedPackageNames"`

	// A list of package ids imported into this instance (if any). Use GetImportedPackage to get the details of the imported packages.
	ImportedPackageIds []string `mandatory:"false" json:"importedPackageIds"`

	// A list of attachment types for this instance (if any). Use attachmentIds to get the details of the attachments.
	AttachmentTypes []string `mandatory:"false" json:"attachmentTypes"`

	// A list of attachment identifiers for this instance (if any). Use GetOdaInstanceAttachment to get the details of the attachments.
	AttachmentIds []string `mandatory:"false" json:"attachmentIds"`

	// A list of restricted operations (across all attachments) for this instance (if any). Use GetOdaInstanceAttachment to get the details of the attachments.
	RestrictedOperations []RestrictedOperation `mandatory:"false" json:"restrictedOperations"`
}

func (m OdaInstance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OdaInstance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOdaInstanceShapeNameEnum(string(m.ShapeName)); !ok && m.ShapeName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ShapeName: %s. Supported values are: %s.", m.ShapeName, strings.Join(GetOdaInstanceShapeNameEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOdaInstanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOdaInstanceLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOdaInstanceLifecycleSubStateEnum(string(m.LifecycleSubState)); !ok && m.LifecycleSubState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleSubState: %s. Supported values are: %s.", m.LifecycleSubState, strings.Join(GetOdaInstanceLifecycleSubStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OdaInstanceShapeNameEnum Enum with underlying type: string
type OdaInstanceShapeNameEnum string

// Set of constants representing the allowable values for OdaInstanceShapeNameEnum
const (
	OdaInstanceShapeNameDevelopment OdaInstanceShapeNameEnum = "DEVELOPMENT"
	OdaInstanceShapeNameProduction  OdaInstanceShapeNameEnum = "PRODUCTION"
)

var mappingOdaInstanceShapeNameEnum = map[string]OdaInstanceShapeNameEnum{
	"DEVELOPMENT": OdaInstanceShapeNameDevelopment,
	"PRODUCTION":  OdaInstanceShapeNameProduction,
}

var mappingOdaInstanceShapeNameEnumLowerCase = map[string]OdaInstanceShapeNameEnum{
	"development": OdaInstanceShapeNameDevelopment,
	"production":  OdaInstanceShapeNameProduction,
}

// GetOdaInstanceShapeNameEnumValues Enumerates the set of values for OdaInstanceShapeNameEnum
func GetOdaInstanceShapeNameEnumValues() []OdaInstanceShapeNameEnum {
	values := make([]OdaInstanceShapeNameEnum, 0)
	for _, v := range mappingOdaInstanceShapeNameEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaInstanceShapeNameEnumStringValues Enumerates the set of values in String for OdaInstanceShapeNameEnum
func GetOdaInstanceShapeNameEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"PRODUCTION",
	}
}

// GetMappingOdaInstanceShapeNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaInstanceShapeNameEnum(val string) (OdaInstanceShapeNameEnum, bool) {
	enum, ok := mappingOdaInstanceShapeNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OdaInstanceLifecycleStateEnum Enum with underlying type: string
type OdaInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for OdaInstanceLifecycleStateEnum
const (
	OdaInstanceLifecycleStateCreating OdaInstanceLifecycleStateEnum = "CREATING"
	OdaInstanceLifecycleStateUpdating OdaInstanceLifecycleStateEnum = "UPDATING"
	OdaInstanceLifecycleStateActive   OdaInstanceLifecycleStateEnum = "ACTIVE"
	OdaInstanceLifecycleStateInactive OdaInstanceLifecycleStateEnum = "INACTIVE"
	OdaInstanceLifecycleStateDeleting OdaInstanceLifecycleStateEnum = "DELETING"
	OdaInstanceLifecycleStateDeleted  OdaInstanceLifecycleStateEnum = "DELETED"
	OdaInstanceLifecycleStateFailed   OdaInstanceLifecycleStateEnum = "FAILED"
)

var mappingOdaInstanceLifecycleStateEnum = map[string]OdaInstanceLifecycleStateEnum{
	"CREATING": OdaInstanceLifecycleStateCreating,
	"UPDATING": OdaInstanceLifecycleStateUpdating,
	"ACTIVE":   OdaInstanceLifecycleStateActive,
	"INACTIVE": OdaInstanceLifecycleStateInactive,
	"DELETING": OdaInstanceLifecycleStateDeleting,
	"DELETED":  OdaInstanceLifecycleStateDeleted,
	"FAILED":   OdaInstanceLifecycleStateFailed,
}

var mappingOdaInstanceLifecycleStateEnumLowerCase = map[string]OdaInstanceLifecycleStateEnum{
	"creating": OdaInstanceLifecycleStateCreating,
	"updating": OdaInstanceLifecycleStateUpdating,
	"active":   OdaInstanceLifecycleStateActive,
	"inactive": OdaInstanceLifecycleStateInactive,
	"deleting": OdaInstanceLifecycleStateDeleting,
	"deleted":  OdaInstanceLifecycleStateDeleted,
	"failed":   OdaInstanceLifecycleStateFailed,
}

// GetOdaInstanceLifecycleStateEnumValues Enumerates the set of values for OdaInstanceLifecycleStateEnum
func GetOdaInstanceLifecycleStateEnumValues() []OdaInstanceLifecycleStateEnum {
	values := make([]OdaInstanceLifecycleStateEnum, 0)
	for _, v := range mappingOdaInstanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaInstanceLifecycleStateEnumStringValues Enumerates the set of values in String for OdaInstanceLifecycleStateEnum
func GetOdaInstanceLifecycleStateEnumStringValues() []string {
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

// GetMappingOdaInstanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaInstanceLifecycleStateEnum(val string) (OdaInstanceLifecycleStateEnum, bool) {
	enum, ok := mappingOdaInstanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OdaInstanceLifecycleSubStateEnum Enum with underlying type: string
type OdaInstanceLifecycleSubStateEnum string

// Set of constants representing the allowable values for OdaInstanceLifecycleSubStateEnum
const (
	OdaInstanceLifecycleSubStateCreating                          OdaInstanceLifecycleSubStateEnum = "CREATING"
	OdaInstanceLifecycleSubStateStarting                          OdaInstanceLifecycleSubStateEnum = "STARTING"
	OdaInstanceLifecycleSubStateStopping                          OdaInstanceLifecycleSubStateEnum = "STOPPING"
	OdaInstanceLifecycleSubStateChangingCompartment               OdaInstanceLifecycleSubStateEnum = "CHANGING_COMPARTMENT"
	OdaInstanceLifecycleSubStateActivatingCustomerEncryptionKey   OdaInstanceLifecycleSubStateEnum = "ACTIVATING_CUSTOMER_ENCRYPTION_KEY"
	OdaInstanceLifecycleSubStateUpdatingCustomerEncryptionKey     OdaInstanceLifecycleSubStateEnum = "UPDATING_CUSTOMER_ENCRYPTION_KEY"
	OdaInstanceLifecycleSubStateDeactivatingCustomerEncryptionKey OdaInstanceLifecycleSubStateEnum = "DEACTIVATING_CUSTOMER_ENCRYPTION_KEY"
	OdaInstanceLifecycleSubStateDeleting                          OdaInstanceLifecycleSubStateEnum = "DELETING"
	OdaInstanceLifecycleSubStateDeletePending                     OdaInstanceLifecycleSubStateEnum = "DELETE_PENDING"
	OdaInstanceLifecycleSubStateRecovering                        OdaInstanceLifecycleSubStateEnum = "RECOVERING"
	OdaInstanceLifecycleSubStateUpdating                          OdaInstanceLifecycleSubStateEnum = "UPDATING"
	OdaInstanceLifecycleSubStatePurging                           OdaInstanceLifecycleSubStateEnum = "PURGING"
	OdaInstanceLifecycleSubStateQueued                            OdaInstanceLifecycleSubStateEnum = "QUEUED"
)

var mappingOdaInstanceLifecycleSubStateEnum = map[string]OdaInstanceLifecycleSubStateEnum{
	"CREATING":                             OdaInstanceLifecycleSubStateCreating,
	"STARTING":                             OdaInstanceLifecycleSubStateStarting,
	"STOPPING":                             OdaInstanceLifecycleSubStateStopping,
	"CHANGING_COMPARTMENT":                 OdaInstanceLifecycleSubStateChangingCompartment,
	"ACTIVATING_CUSTOMER_ENCRYPTION_KEY":   OdaInstanceLifecycleSubStateActivatingCustomerEncryptionKey,
	"UPDATING_CUSTOMER_ENCRYPTION_KEY":     OdaInstanceLifecycleSubStateUpdatingCustomerEncryptionKey,
	"DEACTIVATING_CUSTOMER_ENCRYPTION_KEY": OdaInstanceLifecycleSubStateDeactivatingCustomerEncryptionKey,
	"DELETING":                             OdaInstanceLifecycleSubStateDeleting,
	"DELETE_PENDING":                       OdaInstanceLifecycleSubStateDeletePending,
	"RECOVERING":                           OdaInstanceLifecycleSubStateRecovering,
	"UPDATING":                             OdaInstanceLifecycleSubStateUpdating,
	"PURGING":                              OdaInstanceLifecycleSubStatePurging,
	"QUEUED":                               OdaInstanceLifecycleSubStateQueued,
}

var mappingOdaInstanceLifecycleSubStateEnumLowerCase = map[string]OdaInstanceLifecycleSubStateEnum{
	"creating":                             OdaInstanceLifecycleSubStateCreating,
	"starting":                             OdaInstanceLifecycleSubStateStarting,
	"stopping":                             OdaInstanceLifecycleSubStateStopping,
	"changing_compartment":                 OdaInstanceLifecycleSubStateChangingCompartment,
	"activating_customer_encryption_key":   OdaInstanceLifecycleSubStateActivatingCustomerEncryptionKey,
	"updating_customer_encryption_key":     OdaInstanceLifecycleSubStateUpdatingCustomerEncryptionKey,
	"deactivating_customer_encryption_key": OdaInstanceLifecycleSubStateDeactivatingCustomerEncryptionKey,
	"deleting":                             OdaInstanceLifecycleSubStateDeleting,
	"delete_pending":                       OdaInstanceLifecycleSubStateDeletePending,
	"recovering":                           OdaInstanceLifecycleSubStateRecovering,
	"updating":                             OdaInstanceLifecycleSubStateUpdating,
	"purging":                              OdaInstanceLifecycleSubStatePurging,
	"queued":                               OdaInstanceLifecycleSubStateQueued,
}

// GetOdaInstanceLifecycleSubStateEnumValues Enumerates the set of values for OdaInstanceLifecycleSubStateEnum
func GetOdaInstanceLifecycleSubStateEnumValues() []OdaInstanceLifecycleSubStateEnum {
	values := make([]OdaInstanceLifecycleSubStateEnum, 0)
	for _, v := range mappingOdaInstanceLifecycleSubStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOdaInstanceLifecycleSubStateEnumStringValues Enumerates the set of values in String for OdaInstanceLifecycleSubStateEnum
func GetOdaInstanceLifecycleSubStateEnumStringValues() []string {
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

// GetMappingOdaInstanceLifecycleSubStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOdaInstanceLifecycleSubStateEnum(val string) (OdaInstanceLifecycleSubStateEnum, bool) {
	enum, ok := mappingOdaInstanceLifecycleSubStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
