// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.cloud.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmailTrackConfig The properties that define an email tracking configuration resource.
// Email tracking configuration consists of open tracking, click tracking and list-unsubscribe header.
type EmailTrackConfig struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email tracking configuration resource.
	Id *string `mandatory:"true" json:"id"`

	// Email tracking configuration resource display name.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the email domain resource to which this email tracking configuration applies.
	TrackConfigScopeId *string `mandatory:"true" json:"trackConfigScopeId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Indicates if email open tracking is enabled.
	IsOpenTrackingEnabled *bool `mandatory:"true" json:"isOpenTrackingEnabled"`

	// Indicates if email click tracking is enabled.
	IsClickTrackingEnabled *bool `mandatory:"true" json:"isClickTrackingEnabled"`

	// Indicates if email list unsubscribe header addition is enabled.
	IsListUnsubscribeTrackingEnabled *bool `mandatory:"true" json:"isListUnsubscribeTrackingEnabled"`

	// The current state of the email tracking configuration resource.
	LifecycleState EmailTrackConfigLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The description of an email tracking configuration resource.
	Description *string `mandatory:"false" json:"description"`

	// The resource type to which this email tracking configuration applies.
	TrackConfigScope EmailTrackConfigTrackConfigScopeEnum `mandatory:"false" json:"trackConfigScope,omitempty"`

	// An Internationalized Domain Name used for open and click tracking links once provisioned.
	CustomTrackingDomain *string `mandatory:"false" json:"customTrackingDomain"`

	// The current status of custom tracking domain CNAME setup.
	CustomTrackingDomainStatus CustomTrackingDomainStatusEnum `mandatory:"false" json:"customTrackingDomainStatus,omitempty"`

	// The DNS CNAME record value to be provisioned to the custom tracking domain.
	CustomTrackingDomainCnameRecordValue *string `mandatory:"false" json:"customTrackingDomainCnameRecordValue"`

	// The time the email tracking configuration resource was created. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the email tracking configuration resource was updated. An RFC3339 (https://tools.ietf.org/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m EmailTrackConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmailTrackConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEmailTrackConfigLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEmailTrackConfigLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingEmailTrackConfigTrackConfigScopeEnum(string(m.TrackConfigScope)); !ok && m.TrackConfigScope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TrackConfigScope: %s. Supported values are: %s.", m.TrackConfigScope, strings.Join(GetEmailTrackConfigTrackConfigScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCustomTrackingDomainStatusEnum(string(m.CustomTrackingDomainStatus)); !ok && m.CustomTrackingDomainStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CustomTrackingDomainStatus: %s. Supported values are: %s.", m.CustomTrackingDomainStatus, strings.Join(GetCustomTrackingDomainStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EmailTrackConfigTrackConfigScopeEnum Enum with underlying type: string
type EmailTrackConfigTrackConfigScopeEnum string

// Set of constants representing the allowable values for EmailTrackConfigTrackConfigScopeEnum
const (
	EmailTrackConfigTrackConfigScopeEmaildomain EmailTrackConfigTrackConfigScopeEnum = "EMAILDOMAIN"
)

var mappingEmailTrackConfigTrackConfigScopeEnum = map[string]EmailTrackConfigTrackConfigScopeEnum{
	"EMAILDOMAIN": EmailTrackConfigTrackConfigScopeEmaildomain,
}

var mappingEmailTrackConfigTrackConfigScopeEnumLowerCase = map[string]EmailTrackConfigTrackConfigScopeEnum{
	"emaildomain": EmailTrackConfigTrackConfigScopeEmaildomain,
}

// GetEmailTrackConfigTrackConfigScopeEnumValues Enumerates the set of values for EmailTrackConfigTrackConfigScopeEnum
func GetEmailTrackConfigTrackConfigScopeEnumValues() []EmailTrackConfigTrackConfigScopeEnum {
	values := make([]EmailTrackConfigTrackConfigScopeEnum, 0)
	for _, v := range mappingEmailTrackConfigTrackConfigScopeEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailTrackConfigTrackConfigScopeEnumStringValues Enumerates the set of values in String for EmailTrackConfigTrackConfigScopeEnum
func GetEmailTrackConfigTrackConfigScopeEnumStringValues() []string {
	return []string{
		"EMAILDOMAIN",
	}
}

// GetMappingEmailTrackConfigTrackConfigScopeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailTrackConfigTrackConfigScopeEnum(val string) (EmailTrackConfigTrackConfigScopeEnum, bool) {
	enum, ok := mappingEmailTrackConfigTrackConfigScopeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// EmailTrackConfigLifecycleStateEnum Enum with underlying type: string
type EmailTrackConfigLifecycleStateEnum string

// Set of constants representing the allowable values for EmailTrackConfigLifecycleStateEnum
const (
	EmailTrackConfigLifecycleStateActive   EmailTrackConfigLifecycleStateEnum = "ACTIVE"
	EmailTrackConfigLifecycleStateCreating EmailTrackConfigLifecycleStateEnum = "CREATING"
	EmailTrackConfigLifecycleStateDeleting EmailTrackConfigLifecycleStateEnum = "DELETING"
	EmailTrackConfigLifecycleStateDeleted  EmailTrackConfigLifecycleStateEnum = "DELETED"
	EmailTrackConfigLifecycleStateFailed   EmailTrackConfigLifecycleStateEnum = "FAILED"
	EmailTrackConfigLifecycleStateUpdating EmailTrackConfigLifecycleStateEnum = "UPDATING"
)

var mappingEmailTrackConfigLifecycleStateEnum = map[string]EmailTrackConfigLifecycleStateEnum{
	"ACTIVE":   EmailTrackConfigLifecycleStateActive,
	"CREATING": EmailTrackConfigLifecycleStateCreating,
	"DELETING": EmailTrackConfigLifecycleStateDeleting,
	"DELETED":  EmailTrackConfigLifecycleStateDeleted,
	"FAILED":   EmailTrackConfigLifecycleStateFailed,
	"UPDATING": EmailTrackConfigLifecycleStateUpdating,
}

var mappingEmailTrackConfigLifecycleStateEnumLowerCase = map[string]EmailTrackConfigLifecycleStateEnum{
	"active":   EmailTrackConfigLifecycleStateActive,
	"creating": EmailTrackConfigLifecycleStateCreating,
	"deleting": EmailTrackConfigLifecycleStateDeleting,
	"deleted":  EmailTrackConfigLifecycleStateDeleted,
	"failed":   EmailTrackConfigLifecycleStateFailed,
	"updating": EmailTrackConfigLifecycleStateUpdating,
}

// GetEmailTrackConfigLifecycleStateEnumValues Enumerates the set of values for EmailTrackConfigLifecycleStateEnum
func GetEmailTrackConfigLifecycleStateEnumValues() []EmailTrackConfigLifecycleStateEnum {
	values := make([]EmailTrackConfigLifecycleStateEnum, 0)
	for _, v := range mappingEmailTrackConfigLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEmailTrackConfigLifecycleStateEnumStringValues Enumerates the set of values in String for EmailTrackConfigLifecycleStateEnum
func GetEmailTrackConfigLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"CREATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingEmailTrackConfigLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEmailTrackConfigLifecycleStateEnum(val string) (EmailTrackConfigLifecycleStateEnum, bool) {
	enum, ok := mappingEmailTrackConfigLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
