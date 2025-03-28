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

// ClassicSubscription Subscription information for the compartmentId. Only root compartments are allowed.
type ClassicSubscription struct {

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

	// Classic subscription ID.
	ClassicSubscriptionId *string `mandatory:"true" json:"classicSubscriptionId"`

	// Specifies whether or not the subscription is from classic systems.
	IsClassicSubscription *bool `mandatory:"false" json:"isClassicSubscription"`

	// The pay model of the subscription, such as 'Pay as you go' or 'Monthly'.
	PaymentModel *string `mandatory:"false" json:"paymentModel"`

	// Region for the subscription.
	RegionAssignment *string `mandatory:"false" json:"regionAssignment"`

	// List of SKUs linked to this subscription.
	Skus []SubscriptionSku `mandatory:"false" json:"skus"`

	// Specifies any program that is associated with the subscription.
	ProgramType *string `mandatory:"false" json:"programType"`

	// The country code for the customer associated with the subscription.
	CustomerCountryCode *string `mandatory:"false" json:"customerCountryCode"`

	// The currency code for the customer associated with the subscription.
	CloudAmountCurrency *string `mandatory:"false" json:"cloudAmountCurrency"`

	// Customer service identifier for the customer associated with the subscription.
	CsiNumber *string `mandatory:"false" json:"csiNumber"`

	// Tier for the subscription, whether a free promotion subscription or a paid subscription.
	SubscriptionTier *string `mandatory:"false" json:"subscriptionTier"`

	// Specifies whether or not the subscription is a government subscription.
	IsGovernmentSubscription *bool `mandatory:"false" json:"isGovernmentSubscription"`

	// List of promotions related to the subscription.
	Promotion []Promotion `mandatory:"false" json:"promotion"`

	// Purchase entitlement ID associated with the subscription.
	PurchaseEntitlementId *string `mandatory:"false" json:"purchaseEntitlementId"`

	// Subscription start time.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// Subscription end time.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`

	// This list will contain up to 5 regions where the subscription has a physical presence
	RegionAssignments []string `mandatory:"false" json:"regionAssignments"`

	// Lifecycle state of the subscription.
	LifecycleState ClassicSubscriptionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m ClassicSubscription) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m ClassicSubscription) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetServiceName returns ServiceName
func (m ClassicSubscription) GetServiceName() *string {
	return m.ServiceName
}

// GetTimeCreated returns TimeCreated
func (m ClassicSubscription) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m ClassicSubscription) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m ClassicSubscription) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m ClassicSubscription) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m ClassicSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ClassicSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingClassicSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetClassicSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ClassicSubscription) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeClassicSubscription ClassicSubscription
	s := struct {
		DiscriminatorParam string `json:"entityVersion"`
		MarshalTypeClassicSubscription
	}{
		"V1",
		(MarshalTypeClassicSubscription)(m),
	}

	return json.Marshal(&s)
}
