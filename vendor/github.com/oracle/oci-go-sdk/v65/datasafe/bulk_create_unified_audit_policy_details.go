// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BulkCreateUnifiedAuditPolicyDetails The details required to bulk create unified audit policies.
type BulkCreateUnifiedAuditPolicyDetails struct {

	// The OCID of the security policy corresponding to the unified audit policy.
	SecurityPolicyId *string `mandatory:"true" json:"securityPolicyId"`

	// The OCID of the target database.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The OCID of the compartment in which to create the unified audit policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The list of unified audit policy definition ocids.
	// If unified audit policy definition ids are provided, the imported
	// audit policy will be associated to the specified unified audit policy
	// definition based on the policy name.
	// Else, for every audit policy that gets imported,
	// a new unified audit policy definition will be created.
	UnifiedAuditPolicyDefinitionIds []string `mandatory:"false" json:"unifiedAuditPolicyDefinitionIds"`

	// The list of unified audit policy names to be imported.
	PolicyNames []string `mandatory:"false" json:"policyNames"`

	// Indicates whether the casing of the policy names provided in the request payload should be preserved during creation.
	// By default all policy names will be converted to upper case.
	ShouldPreserveCasing *bool `mandatory:"false" json:"shouldPreserveCasing"`
}

func (m BulkCreateUnifiedAuditPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BulkCreateUnifiedAuditPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
