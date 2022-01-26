// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// EmManagedExternalExadataInsightSummary Summary of an Exadata insight resource.
type EmManagedExternalExadataInsightSummary struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Exadata system name. If the Exadata systems managed by Enterprise Manager, the name is unique amongst the Exadata systems managed by the same Enterprise Manager.
	ExadataName *string `mandatory:"true" json:"exadataName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The time the the Exadata insight was first enabled. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Enterprise Manager Unique Identifier
	EnterpriseManagerIdentifier *string `mandatory:"true" json:"enterpriseManagerIdentifier"`

	// Enterprise Manager Entity Name
	EnterpriseManagerEntityName *string `mandatory:"true" json:"enterpriseManagerEntityName"`

	// Enterprise Manager Entity Type
	EnterpriseManagerEntityType *string `mandatory:"true" json:"enterpriseManagerEntityType"`

	// Enterprise Manager Entity Unique Identifier
	EnterpriseManagerEntityIdentifier *string `mandatory:"true" json:"enterpriseManagerEntityIdentifier"`

	// OPSI Enterprise Manager Bridge OCID
	EnterpriseManagerBridgeId *string `mandatory:"true" json:"enterpriseManagerBridgeId"`

	// The user-friendly name for the Exadata system. The name does not have to be unique.
	ExadataDisplayName *string `mandatory:"false" json:"exadataDisplayName"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the Exadata insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Enterprise Manager Entity Display Name
	EnterpriseManagerEntityDisplayName *string `mandatory:"false" json:"enterpriseManagerEntityDisplayName"`

	// Operations Insights internal representation of the the Exadata system type.
	ExadataType ExadataTypeEnum `mandatory:"false" json:"exadataType,omitempty"`

	// Operations Insights internal representation of the the Exadata system rack type.
	ExadataRackType ExadataRackTypeEnum `mandatory:"false" json:"exadataRackType,omitempty"`

	// Indicates the status of an Exadata insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"true" json:"status"`

	// The current state of the Exadata insight.
	LifecycleState ExadataInsightLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

//GetId returns Id
func (m EmManagedExternalExadataInsightSummary) GetId() *string {
	return m.Id
}

//GetCompartmentId returns CompartmentId
func (m EmManagedExternalExadataInsightSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetExadataName returns ExadataName
func (m EmManagedExternalExadataInsightSummary) GetExadataName() *string {
	return m.ExadataName
}

//GetExadataDisplayName returns ExadataDisplayName
func (m EmManagedExternalExadataInsightSummary) GetExadataDisplayName() *string {
	return m.ExadataDisplayName
}

//GetExadataType returns ExadataType
func (m EmManagedExternalExadataInsightSummary) GetExadataType() ExadataTypeEnum {
	return m.ExadataType
}

//GetExadataRackType returns ExadataRackType
func (m EmManagedExternalExadataInsightSummary) GetExadataRackType() ExadataRackTypeEnum {
	return m.ExadataRackType
}

//GetFreeformTags returns FreeformTags
func (m EmManagedExternalExadataInsightSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m EmManagedExternalExadataInsightSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m EmManagedExternalExadataInsightSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

//GetStatus returns Status
func (m EmManagedExternalExadataInsightSummary) GetStatus() ResourceStatusEnum {
	return m.Status
}

//GetTimeCreated returns TimeCreated
func (m EmManagedExternalExadataInsightSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m EmManagedExternalExadataInsightSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m EmManagedExternalExadataInsightSummary) GetLifecycleState() ExadataInsightLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m EmManagedExternalExadataInsightSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m EmManagedExternalExadataInsightSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m EmManagedExternalExadataInsightSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmManagedExternalExadataInsightSummary EmManagedExternalExadataInsightSummary
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeEmManagedExternalExadataInsightSummary
	}{
		"EM_MANAGED_EXTERNAL_EXADATA",
		(MarshalTypeEmManagedExternalExadataInsightSummary)(m),
	}

	return json.Marshal(&s)
}
