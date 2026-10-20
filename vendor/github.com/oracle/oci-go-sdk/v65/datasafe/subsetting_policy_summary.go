// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubsettingPolicySummary Summary of a subsetting policy
type SubsettingPolicySummary struct {

	// The OCID of the subsetting policy
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the subsetting policy
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the subsetting policy was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the subsetting policy was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The current state of the subsetting policy
	LifecycleState SubsettingPolicyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The display name of the subsetting policy
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the subsetting policy
	Description *string `mandatory:"false" json:"description"`

	SchemaSource SchemaSourceForSummary `mandatory:"false" json:"schemaSource"`

	// The OCID of the masking policy associated with this subsetting policy
	MaskingPolicyId *string `mandatory:"false" json:"maskingPolicyId"`

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

func (m SubsettingPolicySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingPolicySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingPolicyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubsettingPolicyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SubsettingPolicySummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName     *string                            `json:"displayName"`
		Description     *string                            `json:"description"`
		SchemaSource    schemasourceforsummary             `json:"schemaSource"`
		MaskingPolicyId *string                            `json:"maskingPolicyId"`
		FreeformTags    map[string]string                  `json:"freeformTags"`
		DefinedTags     map[string]map[string]interface{}  `json:"definedTags"`
		SystemTags      map[string]map[string]interface{}  `json:"systemTags"`
		Id              *string                            `json:"id"`
		CompartmentId   *string                            `json:"compartmentId"`
		TimeCreated     *common.SDKTime                    `json:"timeCreated"`
		TimeUpdated     *common.SDKTime                    `json:"timeUpdated"`
		LifecycleState  SubsettingPolicyLifecycleStateEnum `json:"lifecycleState"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	nn, e = model.SchemaSource.UnmarshalPolymorphicJSON(model.SchemaSource.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SchemaSource = nn.(SchemaSourceForSummary)
	} else {
		m.SchemaSource = nil
	}

	m.MaskingPolicyId = model.MaskingPolicyId

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	return
}
