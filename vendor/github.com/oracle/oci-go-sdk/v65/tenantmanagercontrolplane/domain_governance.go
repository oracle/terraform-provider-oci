// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DomainGovernance The model for a domain governance entity.
type DomainGovernance struct {

	// The OCID of the domain governance entity.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy that owns this domain governance entity.
	OwnerId *string `mandatory:"true" json:"ownerId"`

	// The OCID of the domain associated with this domain governance entity.
	DomainId *string `mandatory:"true" json:"domainId"`

	// Lifecycle state of the domain governance entity.
	LifecycleState DomainGovernanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The ONS topic associated with this domain governance entity.
	OnsTopicId *string `mandatory:"true" json:"onsTopicId"`

	// The ONS subscription associated with this domain governance entity.
	OnsSubscriptionId *string `mandatory:"true" json:"onsSubscriptionId"`

	// Indicates whether governance is enabled for this domain.
	IsGovernanceEnabled *bool `mandatory:"false" json:"isGovernanceEnabled"`

	// Email address to be used to notify the user, and that the ONS subscription will be created with.
	SubscriptionEmail *string `mandatory:"false" json:"subscriptionEmail"`

	// Date-time when this domain governance was created. An RFC 3339-formatted date and time string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Date-time when this domain governance was last updated. An RFC 3339-formatted date and time string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m DomainGovernance) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DomainGovernance) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDomainGovernanceLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDomainGovernanceLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DomainGovernanceLifecycleStateEnum Enum with underlying type: string
type DomainGovernanceLifecycleStateEnum string

// Set of constants representing the allowable values for DomainGovernanceLifecycleStateEnum
const (
	DomainGovernanceLifecycleStateActive   DomainGovernanceLifecycleStateEnum = "ACTIVE"
	DomainGovernanceLifecycleStateInactive DomainGovernanceLifecycleStateEnum = "INACTIVE"
)

var mappingDomainGovernanceLifecycleStateEnum = map[string]DomainGovernanceLifecycleStateEnum{
	"ACTIVE":   DomainGovernanceLifecycleStateActive,
	"INACTIVE": DomainGovernanceLifecycleStateInactive,
}

var mappingDomainGovernanceLifecycleStateEnumLowerCase = map[string]DomainGovernanceLifecycleStateEnum{
	"active":   DomainGovernanceLifecycleStateActive,
	"inactive": DomainGovernanceLifecycleStateInactive,
}

// GetDomainGovernanceLifecycleStateEnumValues Enumerates the set of values for DomainGovernanceLifecycleStateEnum
func GetDomainGovernanceLifecycleStateEnumValues() []DomainGovernanceLifecycleStateEnum {
	values := make([]DomainGovernanceLifecycleStateEnum, 0)
	for _, v := range mappingDomainGovernanceLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDomainGovernanceLifecycleStateEnumStringValues Enumerates the set of values in String for DomainGovernanceLifecycleStateEnum
func GetDomainGovernanceLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingDomainGovernanceLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDomainGovernanceLifecycleStateEnum(val string) (DomainGovernanceLifecycleStateEnum, bool) {
	enum, ok := mappingDomainGovernanceLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
