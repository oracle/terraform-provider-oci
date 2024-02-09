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

// TargetAlertPolicyAssociation The association of the target database to an alert policy.
type TargetAlertPolicyAssociation struct {

	// The OCID of the target-alert policy association.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the policy.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Creation date and time of the alert policy, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Last date and time the alert policy was updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the target-alert policy association.
	LifecycleState AlertPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The display name of the target-alert policy association.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Describes the target-alert policy association.
	Description *string `mandatory:"false" json:"description"`

	// The OCID of the alert policy.
	PolicyId *string `mandatory:"false" json:"policyId"`

	// The OCID of the target on which alert policy is to be applied.
	TargetId *string `mandatory:"false" json:"targetId"`

	// Indicates if the target-alert policy association is enabled or disabled by user.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Details about the current state of the target-alert policy association.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TargetAlertPolicyAssociation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetAlertPolicyAssociation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlertPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
