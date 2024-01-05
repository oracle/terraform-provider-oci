// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// Recommendation A recommended protection rule for a web application. This recommendation can be accepted to apply it to the Web Application Firewall configuration for this policy.
// Use the `POST /waasPolicies/{waasPolicyId}/actions/acceptWafConfigRecommendations` method to accept recommended protection rules.
type Recommendation struct {

	// The unique key for the recommended protection rule.
	Key *string `mandatory:"false" json:"key"`

	// The list of the ModSecurity rule IDs associated with the protection rule.
	// For more information about ModSecurity's open source WAF rules, see Mod Security's documentation (https://www.modsecurity.org/CRS/Documentation/index.html).
	ModSecurityRuleIds []string `mandatory:"false" json:"modSecurityRuleIds"`

	// The name of the recommended protection rule.
	Name *string `mandatory:"false" json:"name"`

	// The description of the recommended protection rule.
	Description *string `mandatory:"false" json:"description"`

	// The list of labels for the recommended protection rule.
	Labels []string `mandatory:"false" json:"labels"`

	// The recommended action to apply to the protection rule.
	RecommendedAction *string `mandatory:"false" json:"recommendedAction"`
}

func (m Recommendation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Recommendation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
