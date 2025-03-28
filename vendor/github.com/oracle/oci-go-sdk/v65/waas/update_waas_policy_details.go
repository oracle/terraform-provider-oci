// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateWaasPolicyDetails Updates the configuration details of a WAAS policy.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateWaasPolicyDetails struct {

	// A user-friendly name for the WAAS policy. The name can be changed and does not need to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An array of additional domains protected by this WAAS policy.
	AdditionalDomains []string `mandatory:"false" json:"additionalDomains"`

	// A map of host to origin for the web application. The key should be a customer friendly name for the host, ex. primary, secondary, etc.
	Origins map[string]Origin `mandatory:"false" json:"origins"`

	// The map of origin groups and their keys used to associate origins to the `wafConfig`. Origin groups allow you to apply weights to groups of origins for load balancing purposes. Origins with higher weights will receive larger proportions of client requests.
	// To add additional origins to your WAAS policy, update the `origins` field of a `UpdateWaasPolicy` request.
	OriginGroups map[string]OriginGroup `mandatory:"false" json:"originGroups"`

	PolicyConfig *PolicyConfig `mandatory:"false" json:"policyConfig"`

	WafConfig *WafConfig `mandatory:"false" json:"wafConfig"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateWaasPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateWaasPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
