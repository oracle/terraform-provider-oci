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

// TargetDatabaseGroupSummary Summary of the target database group used in list operations. Contains essential information without matching criteria.
type TargetDatabaseGroupSummary struct {

	// The OCID for the compartment containing the target database group.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the specified target database group.
	Id *string `mandatory:"true" json:"id"`

	// The name of the target database group.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The lifecycle status of the target database group.
	LifecycleState TargetDatabaseGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Time when the target database group was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time when the target database group was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Time when the members of the target database group were last changed, i.e. the list was refreshed, a target database was added or removed.
	MembershipUpdateTime *common.SDKTime `mandatory:"true" json:"membershipUpdateTime"`

	// The number of target databases in the specified target database group.
	MembershipCount *int `mandatory:"true" json:"membershipCount"`

	// Description of the target database group.
	Description *string `mandatory:"false" json:"description"`

	// Details for the lifecycle status of the target database group.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TargetDatabaseGroupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetDatabaseGroupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTargetDatabaseGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTargetDatabaseGroupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
