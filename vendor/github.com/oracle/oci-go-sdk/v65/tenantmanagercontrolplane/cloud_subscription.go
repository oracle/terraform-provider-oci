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

// CloudSubscription Subscription information for compartment ID. Only root compartments are allowed.
type CloudSubscription struct {

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

	// Unique Oracle Cloud Subscriptions identifier that is immutable on creation.
	SubscriptionNumber *string `mandatory:"true" json:"subscriptionNumber"`

	// Currency code. For example USD, MXN.
	CurrencyCode *string `mandatory:"true" json:"currencyCode"`

	// Lifecycle state of the subscription.
	LifecycleState SubscriptionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m CloudSubscription) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m CloudSubscription) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetServiceName returns ServiceName
func (m CloudSubscription) GetServiceName() *string {
	return m.ServiceName
}

// GetTimeCreated returns TimeCreated
func (m CloudSubscription) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m CloudSubscription) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetFreeformTags returns FreeformTags
func (m CloudSubscription) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CloudSubscription) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CloudSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CloudSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSubscriptionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubscriptionLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CloudSubscription) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCloudSubscription CloudSubscription
	s := struct {
		DiscriminatorParam string `json:"entityVersion"`
		MarshalTypeCloudSubscription
	}{
		"V2",
		(MarshalTypeCloudSubscription)(m),
	}

	return json.Marshal(&s)
}
