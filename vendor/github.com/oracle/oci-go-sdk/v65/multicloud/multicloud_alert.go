// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Multicloud API
//
// Use the Oracle Multicloud API to retrieve resource anchors and network anchors, and the metadata mappings related a Cloud Service Provider. For more information, see Oracle Multicloud Hub (https://docs.oracle.com/iaas/Content/multicloud-hub/home.htm).
//

package multicloud

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MulticloudAlert A multicloud Alert
type MulticloudAlert struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the multicloud alert.
	Id *string `mandatory:"true" json:"id"`

	// Human-readable name of the alert.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Type/category of the alert (e.g. IAM_POLICY_GAP, TAG_INCONSISTENCY).
	AlertType *string `mandatory:"true" json:"alertType"`

	// Severity of the alert.
	Severity MulticloudAlertSeverityEnum `mandatory:"true" json:"severity"`

	// Current acknowledgment status of the alert.
	AlertStatus MulticloudAlertAlertStatusEnum `mandatory:"true" json:"alertStatus"`

	// Timestamp when the alert was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Internal service or function type generating the alert (e.g. ORP, ODBG_NETWORK, BILLING, OBSERVABILITY).
	FunctionType *string `mandatory:"true" json:"functionType"`

	// Description of the alert and its purpose.
	Description *string `mandatory:"false" json:"description"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the multicloud subscription.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Oracle Cloud Infrastructure Subscription Type.
	SubscriptionType SubscriptionTypeEnum `mandatory:"false" json:"subscriptionType,omitempty"`

	// Timestamp when the alert was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the affected resource.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Type of the affected resource (e.g. ADBD).
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Root Compartment The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) (TenantId) associated with the alert.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Source subsystem that generated the alert. (Azure Tag Validation)
	Source *string `mandatory:"false" json:"source"`

	// External or human-friendly alert identifier.
	AlertId *string `mandatory:"false" json:"alertId"`

	// The current state of the Multicloud Network Alert.
	LifecycleState MulticloudAlertLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// OCI region where the alert originated (e.g. us-phoenix-1)
	SourceRegion *string `mandatory:"false" json:"sourceRegion"`

	// Alert-specific contextual parameters.
	AdditionalParameters map[string]string `mandatory:"false" json:"additionalParameters"`
}

func (m MulticloudAlert) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MulticloudAlert) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMulticloudAlertSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetMulticloudAlertSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMulticloudAlertAlertStatusEnum(string(m.AlertStatus)); !ok && m.AlertStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AlertStatus: %s. Supported values are: %s.", m.AlertStatus, strings.Join(GetMulticloudAlertAlertStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSubscriptionTypeEnum(string(m.SubscriptionType)); !ok && m.SubscriptionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubscriptionType: %s. Supported values are: %s.", m.SubscriptionType, strings.Join(GetSubscriptionTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMulticloudAlertLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMulticloudAlertLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MulticloudAlertSeverityEnum Enum with underlying type: string
type MulticloudAlertSeverityEnum string

// Set of constants representing the allowable values for MulticloudAlertSeverityEnum
const (
	MulticloudAlertSeverityLow      MulticloudAlertSeverityEnum = "LOW"
	MulticloudAlertSeverityMedium   MulticloudAlertSeverityEnum = "MEDIUM"
	MulticloudAlertSeverityHigh     MulticloudAlertSeverityEnum = "HIGH"
	MulticloudAlertSeverityCritical MulticloudAlertSeverityEnum = "CRITICAL"
)

var mappingMulticloudAlertSeverityEnum = map[string]MulticloudAlertSeverityEnum{
	"LOW":      MulticloudAlertSeverityLow,
	"MEDIUM":   MulticloudAlertSeverityMedium,
	"HIGH":     MulticloudAlertSeverityHigh,
	"CRITICAL": MulticloudAlertSeverityCritical,
}

var mappingMulticloudAlertSeverityEnumLowerCase = map[string]MulticloudAlertSeverityEnum{
	"low":      MulticloudAlertSeverityLow,
	"medium":   MulticloudAlertSeverityMedium,
	"high":     MulticloudAlertSeverityHigh,
	"critical": MulticloudAlertSeverityCritical,
}

// GetMulticloudAlertSeverityEnumValues Enumerates the set of values for MulticloudAlertSeverityEnum
func GetMulticloudAlertSeverityEnumValues() []MulticloudAlertSeverityEnum {
	values := make([]MulticloudAlertSeverityEnum, 0)
	for _, v := range mappingMulticloudAlertSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetMulticloudAlertSeverityEnumStringValues Enumerates the set of values in String for MulticloudAlertSeverityEnum
func GetMulticloudAlertSeverityEnumStringValues() []string {
	return []string{
		"LOW",
		"MEDIUM",
		"HIGH",
		"CRITICAL",
	}
}

// GetMappingMulticloudAlertSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMulticloudAlertSeverityEnum(val string) (MulticloudAlertSeverityEnum, bool) {
	enum, ok := mappingMulticloudAlertSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MulticloudAlertAlertStatusEnum Enum with underlying type: string
type MulticloudAlertAlertStatusEnum string

// Set of constants representing the allowable values for MulticloudAlertAlertStatusEnum
const (
	MulticloudAlertAlertStatusUnacknowledged MulticloudAlertAlertStatusEnum = "UNACKNOWLEDGED"
	MulticloudAlertAlertStatusAcknowledged   MulticloudAlertAlertStatusEnum = "ACKNOWLEDGED"
	MulticloudAlertAlertStatusResolved       MulticloudAlertAlertStatusEnum = "RESOLVED"
)

var mappingMulticloudAlertAlertStatusEnum = map[string]MulticloudAlertAlertStatusEnum{
	"UNACKNOWLEDGED": MulticloudAlertAlertStatusUnacknowledged,
	"ACKNOWLEDGED":   MulticloudAlertAlertStatusAcknowledged,
	"RESOLVED":       MulticloudAlertAlertStatusResolved,
}

var mappingMulticloudAlertAlertStatusEnumLowerCase = map[string]MulticloudAlertAlertStatusEnum{
	"unacknowledged": MulticloudAlertAlertStatusUnacknowledged,
	"acknowledged":   MulticloudAlertAlertStatusAcknowledged,
	"resolved":       MulticloudAlertAlertStatusResolved,
}

// GetMulticloudAlertAlertStatusEnumValues Enumerates the set of values for MulticloudAlertAlertStatusEnum
func GetMulticloudAlertAlertStatusEnumValues() []MulticloudAlertAlertStatusEnum {
	values := make([]MulticloudAlertAlertStatusEnum, 0)
	for _, v := range mappingMulticloudAlertAlertStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetMulticloudAlertAlertStatusEnumStringValues Enumerates the set of values in String for MulticloudAlertAlertStatusEnum
func GetMulticloudAlertAlertStatusEnumStringValues() []string {
	return []string{
		"UNACKNOWLEDGED",
		"ACKNOWLEDGED",
		"RESOLVED",
	}
}

// GetMappingMulticloudAlertAlertStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMulticloudAlertAlertStatusEnum(val string) (MulticloudAlertAlertStatusEnum, bool) {
	enum, ok := mappingMulticloudAlertAlertStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// MulticloudAlertLifecycleStateEnum Enum with underlying type: string
type MulticloudAlertLifecycleStateEnum string

// Set of constants representing the allowable values for MulticloudAlertLifecycleStateEnum
const (
	MulticloudAlertLifecycleStateCreating MulticloudAlertLifecycleStateEnum = "CREATING"
	MulticloudAlertLifecycleStateUpdating MulticloudAlertLifecycleStateEnum = "UPDATING"
	MulticloudAlertLifecycleStateActive   MulticloudAlertLifecycleStateEnum = "ACTIVE"
	MulticloudAlertLifecycleStateDeleting MulticloudAlertLifecycleStateEnum = "DELETING"
	MulticloudAlertLifecycleStateDeleted  MulticloudAlertLifecycleStateEnum = "DELETED"
	MulticloudAlertLifecycleStateFailed   MulticloudAlertLifecycleStateEnum = "FAILED"
)

var mappingMulticloudAlertLifecycleStateEnum = map[string]MulticloudAlertLifecycleStateEnum{
	"CREATING": MulticloudAlertLifecycleStateCreating,
	"UPDATING": MulticloudAlertLifecycleStateUpdating,
	"ACTIVE":   MulticloudAlertLifecycleStateActive,
	"DELETING": MulticloudAlertLifecycleStateDeleting,
	"DELETED":  MulticloudAlertLifecycleStateDeleted,
	"FAILED":   MulticloudAlertLifecycleStateFailed,
}

var mappingMulticloudAlertLifecycleStateEnumLowerCase = map[string]MulticloudAlertLifecycleStateEnum{
	"creating": MulticloudAlertLifecycleStateCreating,
	"updating": MulticloudAlertLifecycleStateUpdating,
	"active":   MulticloudAlertLifecycleStateActive,
	"deleting": MulticloudAlertLifecycleStateDeleting,
	"deleted":  MulticloudAlertLifecycleStateDeleted,
	"failed":   MulticloudAlertLifecycleStateFailed,
}

// GetMulticloudAlertLifecycleStateEnumValues Enumerates the set of values for MulticloudAlertLifecycleStateEnum
func GetMulticloudAlertLifecycleStateEnumValues() []MulticloudAlertLifecycleStateEnum {
	values := make([]MulticloudAlertLifecycleStateEnum, 0)
	for _, v := range mappingMulticloudAlertLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMulticloudAlertLifecycleStateEnumStringValues Enumerates the set of values in String for MulticloudAlertLifecycleStateEnum
func GetMulticloudAlertLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMulticloudAlertLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMulticloudAlertLifecycleStateEnum(val string) (MulticloudAlertLifecycleStateEnum, bool) {
	enum, ok := mappingMulticloudAlertLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
