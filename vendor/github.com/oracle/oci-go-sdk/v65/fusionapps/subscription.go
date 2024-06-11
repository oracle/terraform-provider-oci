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

// Subscription Subscription information for compartmentId. Only root compartments are allowed.
type Subscription struct {

	// OCID of the subscription details for particular root compartment or tenancy.
	Id *string `mandatory:"true" json:"id"`

	// Subscription id.
	ClassicSubscriptionId *string `mandatory:"true" json:"classicSubscriptionId"`

	// The type of subscription, such as 'CLOUDCM'/'SAAS'/'CRM', etc.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// Stock keeping unit.
	Skus []SubscriptionSku `mandatory:"true" json:"skus"`

	// Lifecycle state of the subscription.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Subscription resource intermediate states.
	LifecycleDetails SubscriptionLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m Subscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Subscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubscriptionLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetSubscriptionLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubscriptionLifecycleStateEnum Enum with underlying type: string
type SubscriptionLifecycleStateEnum string

// Set of constants representing the allowable values for SubscriptionLifecycleStateEnum
const (
	SubscriptionLifecycleStateCreating       SubscriptionLifecycleStateEnum = "CREATING"
	SubscriptionLifecycleStateActive         SubscriptionLifecycleStateEnum = "ACTIVE"
	SubscriptionLifecycleStateInactive       SubscriptionLifecycleStateEnum = "INACTIVE"
	SubscriptionLifecycleStateUpdating       SubscriptionLifecycleStateEnum = "UPDATING"
	SubscriptionLifecycleStateDeleting       SubscriptionLifecycleStateEnum = "DELETING"
	SubscriptionLifecycleStateDeleted        SubscriptionLifecycleStateEnum = "DELETED"
	SubscriptionLifecycleStateFailed         SubscriptionLifecycleStateEnum = "FAILED"
	SubscriptionLifecycleStateNeedsAttention SubscriptionLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingSubscriptionLifecycleStateEnum = map[string]SubscriptionLifecycleStateEnum{
	"CREATING":        SubscriptionLifecycleStateCreating,
	"ACTIVE":          SubscriptionLifecycleStateActive,
	"INACTIVE":        SubscriptionLifecycleStateInactive,
	"UPDATING":        SubscriptionLifecycleStateUpdating,
	"DELETING":        SubscriptionLifecycleStateDeleting,
	"DELETED":         SubscriptionLifecycleStateDeleted,
	"FAILED":          SubscriptionLifecycleStateFailed,
	"NEEDS_ATTENTION": SubscriptionLifecycleStateNeedsAttention,
}

var mappingSubscriptionLifecycleStateEnumLowerCase = map[string]SubscriptionLifecycleStateEnum{
	"creating":        SubscriptionLifecycleStateCreating,
	"active":          SubscriptionLifecycleStateActive,
	"inactive":        SubscriptionLifecycleStateInactive,
	"updating":        SubscriptionLifecycleStateUpdating,
	"deleting":        SubscriptionLifecycleStateDeleting,
	"deleted":         SubscriptionLifecycleStateDeleted,
	"failed":          SubscriptionLifecycleStateFailed,
	"needs_attention": SubscriptionLifecycleStateNeedsAttention,
}

// GetSubscriptionLifecycleStateEnumValues Enumerates the set of values for SubscriptionLifecycleStateEnum
func GetSubscriptionLifecycleStateEnumValues() []SubscriptionLifecycleStateEnum {
	values := make([]SubscriptionLifecycleStateEnum, 0)
	for _, v := range mappingSubscriptionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionLifecycleStateEnumStringValues Enumerates the set of values in String for SubscriptionLifecycleStateEnum
func GetSubscriptionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingSubscriptionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionLifecycleStateEnum(val string) (SubscriptionLifecycleStateEnum, bool) {
	enum, ok := mappingSubscriptionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SubscriptionLifecycleDetailsEnum Enum with underlying type: string
type SubscriptionLifecycleDetailsEnum string

// Set of constants representing the allowable values for SubscriptionLifecycleDetailsEnum
const (
	SubscriptionLifecycleDetailsActive                  SubscriptionLifecycleDetailsEnum = "ACTIVE"
	SubscriptionLifecycleDetailsPurged                  SubscriptionLifecycleDetailsEnum = "PURGED"
	SubscriptionLifecycleDetailsCanceled                SubscriptionLifecycleDetailsEnum = "CANCELED"
	SubscriptionLifecycleDetailsPaidPurged              SubscriptionLifecycleDetailsEnum = "PAID_PURGED"
	SubscriptionLifecycleDetailsInitialized             SubscriptionLifecycleDetailsEnum = "INITIALIZED"
	SubscriptionLifecycleDetailsSoftTerminated          SubscriptionLifecycleDetailsEnum = "SOFT_TERMINATED"
	SubscriptionLifecycleDetailsDisabled                SubscriptionLifecycleDetailsEnum = "DISABLED"
	SubscriptionLifecycleDetailsBeginTermination        SubscriptionLifecycleDetailsEnum = "BEGIN_TERMINATION"
	SubscriptionLifecycleDetailsMigrated                SubscriptionLifecycleDetailsEnum = "MIGRATED"
	SubscriptionLifecycleDetailsPendingCancelation      SubscriptionLifecycleDetailsEnum = "PENDING_CANCELATION"
	SubscriptionLifecycleDetailsArchived                SubscriptionLifecycleDetailsEnum = "ARCHIVED"
	SubscriptionLifecycleDetailsNonRecoverable          SubscriptionLifecycleDetailsEnum = "NON_RECOVERABLE"
	SubscriptionLifecycleDetailsBeginSoftTermination    SubscriptionLifecycleDetailsEnum = "BEGIN_SOFT_TERMINATION"
	SubscriptionLifecycleDetailsActivated               SubscriptionLifecycleDetailsEnum = "ACTIVATED"
	SubscriptionLifecycleDetailsAccessDisabled          SubscriptionLifecycleDetailsEnum = "ACCESS_DISABLED"
	SubscriptionLifecycleDetailsPendingRegistration     SubscriptionLifecycleDetailsEnum = "PENDING_REGISTRATION"
	SubscriptionLifecycleDetailsTerminated              SubscriptionLifecycleDetailsEnum = "TERMINATED"
	SubscriptionLifecycleDetailsRelocating              SubscriptionLifecycleDetailsEnum = "RELOCATING"
	SubscriptionLifecycleDetailsDeprovisioned           SubscriptionLifecycleDetailsEnum = "DEPROVISIONED"
	SubscriptionLifecycleDetailsProvisioned             SubscriptionLifecycleDetailsEnum = "PROVISIONED"
	SubscriptionLifecycleDetailsBeginTerminationPassive SubscriptionLifecycleDetailsEnum = "BEGIN_TERMINATION_PASSIVE"
	SubscriptionLifecycleDetailsLocked                  SubscriptionLifecycleDetailsEnum = "LOCKED"
	SubscriptionLifecycleDetailsPendingDeprovisioning   SubscriptionLifecycleDetailsEnum = "PENDING_DEPROVISIONING"
	SubscriptionLifecycleDetailsRegistered              SubscriptionLifecycleDetailsEnum = "REGISTERED"
	SubscriptionLifecycleDetailsCancelled               SubscriptionLifecycleDetailsEnum = "CANCELLED"
	SubscriptionLifecycleDetailsExpired                 SubscriptionLifecycleDetailsEnum = "EXPIRED"
)

var mappingSubscriptionLifecycleDetailsEnum = map[string]SubscriptionLifecycleDetailsEnum{
	"ACTIVE":                    SubscriptionLifecycleDetailsActive,
	"PURGED":                    SubscriptionLifecycleDetailsPurged,
	"CANCELED":                  SubscriptionLifecycleDetailsCanceled,
	"PAID_PURGED":               SubscriptionLifecycleDetailsPaidPurged,
	"INITIALIZED":               SubscriptionLifecycleDetailsInitialized,
	"SOFT_TERMINATED":           SubscriptionLifecycleDetailsSoftTerminated,
	"DISABLED":                  SubscriptionLifecycleDetailsDisabled,
	"BEGIN_TERMINATION":         SubscriptionLifecycleDetailsBeginTermination,
	"MIGRATED":                  SubscriptionLifecycleDetailsMigrated,
	"PENDING_CANCELATION":       SubscriptionLifecycleDetailsPendingCancelation,
	"ARCHIVED":                  SubscriptionLifecycleDetailsArchived,
	"NON_RECOVERABLE":           SubscriptionLifecycleDetailsNonRecoverable,
	"BEGIN_SOFT_TERMINATION":    SubscriptionLifecycleDetailsBeginSoftTermination,
	"ACTIVATED":                 SubscriptionLifecycleDetailsActivated,
	"ACCESS_DISABLED":           SubscriptionLifecycleDetailsAccessDisabled,
	"PENDING_REGISTRATION":      SubscriptionLifecycleDetailsPendingRegistration,
	"TERMINATED":                SubscriptionLifecycleDetailsTerminated,
	"RELOCATING":                SubscriptionLifecycleDetailsRelocating,
	"DEPROVISIONED":             SubscriptionLifecycleDetailsDeprovisioned,
	"PROVISIONED":               SubscriptionLifecycleDetailsProvisioned,
	"BEGIN_TERMINATION_PASSIVE": SubscriptionLifecycleDetailsBeginTerminationPassive,
	"LOCKED":                    SubscriptionLifecycleDetailsLocked,
	"PENDING_DEPROVISIONING":    SubscriptionLifecycleDetailsPendingDeprovisioning,
	"REGISTERED":                SubscriptionLifecycleDetailsRegistered,
	"CANCELLED":                 SubscriptionLifecycleDetailsCancelled,
	"EXPIRED":                   SubscriptionLifecycleDetailsExpired,
}

var mappingSubscriptionLifecycleDetailsEnumLowerCase = map[string]SubscriptionLifecycleDetailsEnum{
	"active":                    SubscriptionLifecycleDetailsActive,
	"purged":                    SubscriptionLifecycleDetailsPurged,
	"canceled":                  SubscriptionLifecycleDetailsCanceled,
	"paid_purged":               SubscriptionLifecycleDetailsPaidPurged,
	"initialized":               SubscriptionLifecycleDetailsInitialized,
	"soft_terminated":           SubscriptionLifecycleDetailsSoftTerminated,
	"disabled":                  SubscriptionLifecycleDetailsDisabled,
	"begin_termination":         SubscriptionLifecycleDetailsBeginTermination,
	"migrated":                  SubscriptionLifecycleDetailsMigrated,
	"pending_cancelation":       SubscriptionLifecycleDetailsPendingCancelation,
	"archived":                  SubscriptionLifecycleDetailsArchived,
	"non_recoverable":           SubscriptionLifecycleDetailsNonRecoverable,
	"begin_soft_termination":    SubscriptionLifecycleDetailsBeginSoftTermination,
	"activated":                 SubscriptionLifecycleDetailsActivated,
	"access_disabled":           SubscriptionLifecycleDetailsAccessDisabled,
	"pending_registration":      SubscriptionLifecycleDetailsPendingRegistration,
	"terminated":                SubscriptionLifecycleDetailsTerminated,
	"relocating":                SubscriptionLifecycleDetailsRelocating,
	"deprovisioned":             SubscriptionLifecycleDetailsDeprovisioned,
	"provisioned":               SubscriptionLifecycleDetailsProvisioned,
	"begin_termination_passive": SubscriptionLifecycleDetailsBeginTerminationPassive,
	"locked":                    SubscriptionLifecycleDetailsLocked,
	"pending_deprovisioning":    SubscriptionLifecycleDetailsPendingDeprovisioning,
	"registered":                SubscriptionLifecycleDetailsRegistered,
	"cancelled":                 SubscriptionLifecycleDetailsCancelled,
	"expired":                   SubscriptionLifecycleDetailsExpired,
}

// GetSubscriptionLifecycleDetailsEnumValues Enumerates the set of values for SubscriptionLifecycleDetailsEnum
func GetSubscriptionLifecycleDetailsEnumValues() []SubscriptionLifecycleDetailsEnum {
	values := make([]SubscriptionLifecycleDetailsEnum, 0)
	for _, v := range mappingSubscriptionLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetSubscriptionLifecycleDetailsEnumStringValues Enumerates the set of values in String for SubscriptionLifecycleDetailsEnum
func GetSubscriptionLifecycleDetailsEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"PURGED",
		"CANCELED",
		"PAID_PURGED",
		"INITIALIZED",
		"SOFT_TERMINATED",
		"DISABLED",
		"BEGIN_TERMINATION",
		"MIGRATED",
		"PENDING_CANCELATION",
		"ARCHIVED",
		"NON_RECOVERABLE",
		"BEGIN_SOFT_TERMINATION",
		"ACTIVATED",
		"ACCESS_DISABLED",
		"PENDING_REGISTRATION",
		"TERMINATED",
		"RELOCATING",
		"DEPROVISIONED",
		"PROVISIONED",
		"BEGIN_TERMINATION_PASSIVE",
		"LOCKED",
		"PENDING_DEPROVISIONING",
		"REGISTERED",
		"CANCELLED",
		"EXPIRED",
	}
}

// GetMappingSubscriptionLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubscriptionLifecycleDetailsEnum(val string) (SubscriptionLifecycleDetailsEnum, bool) {
	enum, ok := mappingSubscriptionLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
