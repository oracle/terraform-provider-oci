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

// CreateChargebackPlanDetails The details used to create a new Ops Insights chargeback plan.
type CreateChargebackPlanDetails interface {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// Name for the OPSI Chargeback plan.
	GetPlanName() *string

	// Chargeback Plan type of the chargeback entity. For an Exadata it can be WEIGHTED_ALLOCATION, EQUAL_ALLOCATION, UNUSED_ALLOCATION.
	GetPlanType() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Description of OPSI Chargeback Plan.
	GetPlanDescription() *string

	// List of chargeback plan customizations.
	GetPlanCustomItems() []CreatePlanCustomItemDetails
}

type createchargebackplandetails struct {
	JsonData        []byte
	FreeformTags    map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags     map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	PlanDescription *string                           `mandatory:"false" json:"planDescription"`
	PlanCustomItems []CreatePlanCustomItemDetails     `mandatory:"false" json:"planCustomItems"`
	CompartmentId   *string                           `mandatory:"true" json:"compartmentId"`
	PlanName        *string                           `mandatory:"true" json:"planName"`
	PlanType        *string                           `mandatory:"true" json:"planType"`
	EntitySource    string                            `json:"entitySource"`
}

// UnmarshalJSON unmarshals json
func (m *createchargebackplandetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatechargebackplandetails createchargebackplandetails
	s := struct {
		Model Unmarshalercreatechargebackplandetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.PlanName = s.Model.PlanName
	m.PlanType = s.Model.PlanType
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.PlanDescription = s.Model.PlanDescription
	m.PlanCustomItems = s.Model.PlanCustomItems
	m.EntitySource = s.Model.EntitySource

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createchargebackplandetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EntitySource {
	case "CHARGEBACK_EXADATA":
		mm := CreateChargebackPlanExadataDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateChargebackPlanDetails: %s.", m.EntitySource)
		return *m, nil
	}
}

// GetFreeformTags returns FreeformTags
func (m createchargebackplandetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createchargebackplandetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetPlanDescription returns PlanDescription
func (m createchargebackplandetails) GetPlanDescription() *string {
	return m.PlanDescription
}

// GetPlanCustomItems returns PlanCustomItems
func (m createchargebackplandetails) GetPlanCustomItems() []CreatePlanCustomItemDetails {
	return m.PlanCustomItems
}

// GetCompartmentId returns CompartmentId
func (m createchargebackplandetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetPlanName returns PlanName
func (m createchargebackplandetails) GetPlanName() *string {
	return m.PlanName
}

// GetPlanType returns PlanType
func (m createchargebackplandetails) GetPlanType() *string {
	return m.PlanType
}

func (m createchargebackplandetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createchargebackplandetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
