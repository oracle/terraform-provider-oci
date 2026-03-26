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
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateChargebackPlanExadataDetails The information about the exadata and chargeback plan.
type CreateChargebackPlanExadataDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name for the OPSI Chargeback plan.
	PlanName *string `mandatory:"true" json:"planName"`

	// Chargeback Plan type of the chargeback entity. For an Exadata it can be WEIGHTED_ALLOCATION, EQUAL_ALLOCATION, UNUSED_ALLOCATION.
	PlanType *string `mandatory:"true" json:"planType"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Description of OPSI Chargeback Plan.
	PlanDescription *string `mandatory:"false" json:"planDescription"`

	// List of chargeback plan customizations.
	PlanCustomItems []CreatePlanCustomItemDetails `mandatory:"false" json:"planCustomItems"`
}

// GetCompartmentId returns CompartmentId
func (m CreateChargebackPlanExadataDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetFreeformTags returns FreeformTags
func (m CreateChargebackPlanExadataDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateChargebackPlanExadataDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetPlanName returns PlanName
func (m CreateChargebackPlanExadataDetails) GetPlanName() *string {
	return m.PlanName
}

// GetPlanDescription returns PlanDescription
func (m CreateChargebackPlanExadataDetails) GetPlanDescription() *string {
	return m.PlanDescription
}

// GetPlanType returns PlanType
func (m CreateChargebackPlanExadataDetails) GetPlanType() *string {
	return m.PlanType
}

// GetPlanCustomItems returns PlanCustomItems
func (m CreateChargebackPlanExadataDetails) GetPlanCustomItems() []CreatePlanCustomItemDetails {
	return m.PlanCustomItems
}

func (m CreateChargebackPlanExadataDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateChargebackPlanExadataDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateChargebackPlanExadataDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateChargebackPlanExadataDetails CreateChargebackPlanExadataDetails
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeCreateChargebackPlanExadataDetails
	}{
		"CHARGEBACK_EXADATA",
		(MarshalTypeCreateChargebackPlanExadataDetails)(m),
	}

	return json.Marshal(&s)
}
