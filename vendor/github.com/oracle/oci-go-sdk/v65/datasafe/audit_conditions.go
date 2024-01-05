// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// AuditConditions Represents audit policies with corresponding audit provisioning conditions.
type AuditConditions struct {

	// Indicates the audit policy name. Refer to the documentation (https://docs.oracle.com/en/cloud/paas/data-safe/udscs/audit-policies.html#GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827) for seeded audit policy names. For custom policies, refer to the user-defined policy name created in the target database.
	AuditPolicyName *string `mandatory:"true" json:"auditPolicyName"`

	// Indicates whether the privileged user list is managed by Data Safe.
	IsPrivUsersManagedByDataSafe *bool `mandatory:"true" json:"isPrivUsersManagedByDataSafe"`

	// Indicates whether the Data Safe user activity on the target database will be audited by the policy.
	IsDataSafeServiceAccountAudited *bool `mandatory:"true" json:"isDataSafeServiceAccountAudited"`

	// Indicates the users/roles in the target database for which the audit policy is enforced, and the success/failure event condition to generate the audit event..
	EnableConditions []EnableConditions `mandatory:"false" json:"enableConditions"`
}

func (m AuditConditions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditConditions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
