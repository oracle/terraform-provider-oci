// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AuditPolicyDimensions Details of aggregation dimensions used for summarizing audit policies.
type AuditPolicyDimensions struct {

	// The category to which the audit policy belongs.
	AuditPolicyCategory AuditPolicyCategoryEnum `mandatory:"false" json:"auditPolicyCategory,omitempty"`

	// The name of the audit policy. Refer to the documentation (https://www.oracle.com/pls/topic/lookup?ctx=en/cloud/paas/data-safe&id=UDSCS-GUID-361A9A9A-7C21-4F5A-8945-9B3A0C472827) for seeded audit policy names. For custom policies, refer to the user-defined policy name created in the target database.
	AuditPolicyName *string `mandatory:"false" json:"auditPolicyName"`

	// The OCID of the target database for which the audit policy will be created.
	TargetId *string `mandatory:"false" json:"targetId"`
}

func (m AuditPolicyDimensions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuditPolicyDimensions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAuditPolicyCategoryEnum(string(m.AuditPolicyCategory)); !ok && m.AuditPolicyCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AuditPolicyCategory: %s. Supported values are: %s.", m.AuditPolicyCategory, strings.Join(GetAuditPolicyCategoryEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
