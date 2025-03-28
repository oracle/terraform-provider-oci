// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StackSummary Summary information for a stack.
type StackSummary struct {

	// Unique identifier of the specified stack.
	Id *string `mandatory:"false" json:"id"`

	// Unique identifier of the compartment in which the stack resides.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Human-readable display name for the stack.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// General description of the stack.
	Description *string `mandatory:"false" json:"description"`

	// The date and time when the stack was created.
	// Format is defined by RFC3339.
	// Example: `2020-01-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The current lifecycle state of the stack.
	// For more information about stack lifecycle states in Resource Manager, see
	// Key Concepts (https://docs.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#concepts__StackStates).
	// Allowable values:
	// - CREATING
	// - ACTIVE
	// - DELETING
	// - DELETED
	// - FAILED
	LifecycleState StackLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// The version of Terraform specified for the stack. Example: `0.12.x`
	TerraformVersion *string `mandatory:"false" json:"terraformVersion"`

	// Free-form tags associated with this resource. Each tag is a key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m StackSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStackLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetStackLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
