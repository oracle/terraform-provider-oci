// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CustomProtectionRuleSummary Summary information about a Custom Protection rule.
type CustomProtectionRuleSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Custom Protection rule.
	Id *string `mandatory:"false" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the Custom Protection rule's compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The user-friendly name of the Custom Protection rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The list of the ModSecurity rule IDs that apply to this protection rule. For more information about ModSecurity's open source WAF rules, see Mod Security's documentation (https://www.modsecurity.org/CRS/Documentation/index.html).
	ModSecurityRuleIds []string `mandatory:"false" json:"modSecurityRuleIds"`

	// The current lifecycle state of the Custom Protection rule.
	LifecycleState LifecycleStatesEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The date and time the protection rule was created, expressed in RFC 3339 timestamp format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CustomProtectionRuleSummary) String() string {
	return common.PointerString(m)
}
