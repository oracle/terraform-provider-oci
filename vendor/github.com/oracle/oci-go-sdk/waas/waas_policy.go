// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WaasPolicy The details of a Web Application Acceleration and Security (WAAS) policy. A policy describes how the WAAS service should operate for the configured web application.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type WaasPolicy struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WAAS policy.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the WAAS policy's compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user-friendly name of the WAAS policy. The name can be changed and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The web application domain that the WAAS policy protects.
	Domain *string `mandatory:"false" json:"domain"`

	// An array of additional domains for this web application.
	AdditionalDomains []string `mandatory:"false" json:"additionalDomains"`

	// The CNAME record to add to your DNS configuration to route traffic for the domain, and all additional domains, through the WAF.
	Cname *string `mandatory:"false" json:"cname"`

	// The current lifecycle state of the WAAS policy.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the policy was created, expressed in RFC 3339 timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// A map of host servers (origins) and their keys for the web application. Origin keys are used to associate origins to specific protection rules. The key should be a user-friendly name for the host. **Examples:** `primary` or `secondary`.
	Origins map[string]Origin `mandatory:"false" json:"origins"`

	// The map of origin groups and their keys used to associate origins to the `wafConfig`. Origin groups allow you to apply weights to groups of origins for load balancing purposes. Origins with higher weights will receive larger proportions of client requests.
	OriginGroups map[string]OriginGroup `mandatory:"false" json:"originGroups"`

	PolicyConfig *PolicyConfig `mandatory:"false" json:"policyConfig"`

	WafConfig *WafConfig `mandatory:"false" json:"wafConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m WaasPolicy) String() string {
	return common.PointerString(m)
}

// WaasPolicyLifecycleStateEnum is an alias to type: LifecycleStatesEnum
// Consider using LifecycleStatesEnum instead
// Deprecated
type WaasPolicyLifecycleStateEnum = LifecycleStatesEnum

// Set of constants representing the allowable values for LifecycleStatesEnum
// Deprecated
const (
	WaasPolicyLifecycleStateCreating LifecycleStatesEnum = "CREATING"
	WaasPolicyLifecycleStateActive   LifecycleStatesEnum = "ACTIVE"
	WaasPolicyLifecycleStateFailed   LifecycleStatesEnum = "FAILED"
	WaasPolicyLifecycleStateUpdating LifecycleStatesEnum = "UPDATING"
	WaasPolicyLifecycleStateDeleting LifecycleStatesEnum = "DELETING"
	WaasPolicyLifecycleStateDeleted  LifecycleStatesEnum = "DELETED"
)

// GetWaasPolicyLifecycleStateEnumValues Enumerates the set of values for LifecycleStatesEnum
// Consider using GetLifecycleStatesEnumValue
// Deprecated
var GetWaasPolicyLifecycleStateEnumValues = GetLifecycleStatesEnumValues
