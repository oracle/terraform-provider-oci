// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ClassicAssignedSubscriptionSummary Summary of assigned subscription information.
type ClassicAssignedSubscriptionSummary struct {

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the subscription.
	Id *string `mandatory:"true" json:"id"`

	// The Oracle ID (OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm)) of the owning compartment. Always a tenancy OCID.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The type of subscription, such as 'UCM', 'SAAS', 'ERP', 'CRM'.
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The date and time of creation, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time of update, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// Subscription ID.
	ClassicSubscriptionId *string `mandatory:"true" json:"classicSubscriptionId"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Specifies whether or not the subscription is legacy.
	IsClassicSubscription *bool `mandatory:"false" json:"isClassicSubscription"`

	// Region for the subscription.
	RegionAssignment *string `mandatory:"false" json:"regionAssignment"`

	// Subscription start time.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// Subscription end time.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`

	// Customer service identifier for the customer associated with the subscription.
	CsiNumber *string `mandatory:"false" json:"csiNumber"`

	// Lifecycle state of the subscription.
	LifecycleState ClassicSubscriptionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Service or component which is used to provision and manage the subscription.
	ManagedBy ClassicSubscriptionManagedByEnum `mandatory:"false" json:"managedBy,omitempty"`
}

// GetId returns Id
func (m ClassicAssignedSubscriptionSummary) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m ClassicAssignedSubscriptionSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetServiceName returns ServiceName
func (m ClassicAssignedSubscriptionSummary) GetServiceName() *string {
	return m.ServiceName
}

// GetTimeCreated returns TimeCreated
func (m ClassicAssignedSubscriptionSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ClassicAssignedSubscriptionSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m ClassicAssignedSubscriptionSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m ClassicAssignedSubscriptionSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m ClassicAssignedSubscriptionSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m ClassicAssignedSubscriptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClassicAssignedSubscriptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClassicSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetClassicSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingClassicSubscriptionManagedByEnum(string(m.ManagedBy)); !ok && m.ManagedBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedBy: %s. Supported values are: %s.", m.ManagedBy, strings.Join(GetClassicSubscriptionManagedByEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ClassicAssignedSubscriptionSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeClassicAssignedSubscriptionSummary ClassicAssignedSubscriptionSummary
	s := struct {
		DiscriminatorParam string `json:"entityVersion"`
		MarshalTypeClassicAssignedSubscriptionSummary
	}{
		"V1",
		(MarshalTypeClassicAssignedSubscriptionSummary)(m),
	}

	return json.Marshal(&s)
}
