// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ChargebackPlan A chargeback plan that allows Ops Insights services to compute chargeback costs.
type ChargebackPlan struct {

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OPSI Chargeback plan resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name for the OPSI Chargeback plan.
	PlanName *string `mandatory:"true" json:"planName"`

	// Chargeback Plan type of the chargeback entity. For an Exadata it can be WEIGHTED_ALLOCATION, EQUAL_ALLOCATION, UNUSED_ALLOCATION.
	PlanType *string `mandatory:"true" json:"planType"`

	// The date and time the chargeback plan was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Chargeback Plan lifecycle states
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Description of OPSI Chargeback Plan.
	PlanDescription *string `mandatory:"false" json:"planDescription"`

	// Chargeback Plan category of the chargeback entity. It can be OOB, or CUSTOM.
	PlanCategory ChargebackPlanCategoryEnum `mandatory:"false" json:"planCategory,omitempty"`

	// Indicates whether the chargeback plan can be customized.
	IsCustomizable *bool `mandatory:"false" json:"isCustomizable"`

	// Source of the chargeback plan.
	EntitySource ChargebackPlanEntitySourceEnum `mandatory:"false" json:"entitySource,omitempty"`

	// The time chargeback plan was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// List of chargeback plan customizations.
	PlanCustomItems []CreatePlanCustomItemDetails `mandatory:"false" json:"planCustomItems"`
}

func (m ChargebackPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ChargebackPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingChargebackPlanCategoryEnum(string(m.PlanCategory)); !ok && m.PlanCategory != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanCategory: %s. Supported values are: %s.", m.PlanCategory, strings.Join(GetChargebackPlanCategoryEnumStringValues(), ",")))
	}
	if _, ok := GetMappingChargebackPlanEntitySourceEnum(string(m.EntitySource)); !ok && m.EntitySource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EntitySource: %s. Supported values are: %s.", m.EntitySource, strings.Join(GetChargebackPlanEntitySourceEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
