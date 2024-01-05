// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EmManagedExternalExadataInsight EM-managed Exadata insight resource.
type EmManagedExternalExadataInsight struct {

	// Exadata insight identifier
	Id *string `mandatory:"true" json:"id"`

	// Compartment identifier of the Exadata insight resource
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

	// true if virtualization is used in the Exadata system
	IsVirtualizedExadata *bool `mandatory:"false" json:"isVirtualizedExadata"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The time the Exadata insight was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Enterprise Manager Entity Display Name
	EnterpriseManagerEntityDisplayName *string `mandatory:"false" json:"enterpriseManagerEntityDisplayName"`

	// Set to true to enable automatic enablement and disablement of related targets from Enterprise Manager. New resources (e.g. Database Insights) will be placed in the same compartment as the related Exadata Insight.
	IsAutoSyncEnabled *bool `mandatory:"false" json:"isAutoSyncEnabled"`

	// Operations Insights internal representation of the the Exadata system type.
	ExadataType ExadataTypeEnum `mandatory:"false" json:"exadataType,omitempty"`

	// Exadata rack type.
	ExadataRackType ExadataRackTypeEnum `mandatory:"false" json:"exadataRackType,omitempty"`

	// Indicates the status of an Exadata insight in Operations Insights
	Status ResourceStatusEnum `mandatory:"true" json:"status"`

	// The current state of the Exadata insight.
	LifecycleState ExadataInsightLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m EmManagedExternalExadataInsight) GetId() *string {
	return m.Id
}

// GetCompartmentId returns CompartmentId
func (m EmManagedExternalExadataInsight) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetExadataName returns ExadataName
func (m EmManagedExternalExadataInsight) GetExadataName() *string {
	return m.ExadataName
}

// GetExadataDisplayName returns ExadataDisplayName
func (m EmManagedExternalExadataInsight) GetExadataDisplayName() *string {
	return m.ExadataDisplayName
}

// GetExadataType returns ExadataType
func (m EmManagedExternalExadataInsight) GetExadataType() ExadataTypeEnum {
	return m.ExadataType
}

// GetExadataRackType returns ExadataRackType
func (m EmManagedExternalExadataInsight) GetExadataRackType() ExadataRackTypeEnum {
	return m.ExadataRackType
}

// GetIsVirtualizedExadata returns IsVirtualizedExadata
func (m EmManagedExternalExadataInsight) GetIsVirtualizedExadata() *bool {
	return m.IsVirtualizedExadata
}

// GetStatus returns Status
func (m EmManagedExternalExadataInsight) GetStatus() ResourceStatusEnum {
	return m.Status
}

// GetFreeformTags returns FreeformTags
func (m EmManagedExternalExadataInsight) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m EmManagedExternalExadataInsight) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m EmManagedExternalExadataInsight) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetTimeCreated returns TimeCreated
func (m EmManagedExternalExadataInsight) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m EmManagedExternalExadataInsight) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m EmManagedExternalExadataInsight) GetLifecycleState() ExadataInsightLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m EmManagedExternalExadataInsight) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

func (m EmManagedExternalExadataInsight) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EmManagedExternalExadataInsight) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingExadataTypeEnum(string(m.ExadataType)); !ok && m.ExadataType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataType: %s. Supported values are: %s.", m.ExadataType, strings.Join(GetExadataTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataRackTypeEnum(string(m.ExadataRackType)); !ok && m.ExadataRackType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ExadataRackType: %s. Supported values are: %s.", m.ExadataRackType, strings.Join(GetExadataRackTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetResourceStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingExadataInsightLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetExadataInsightLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m EmManagedExternalExadataInsight) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeEmManagedExternalExadataInsight EmManagedExternalExadataInsight
	s := struct {
		DiscriminatorParam string `json:"entitySource"`
		MarshalTypeEmManagedExternalExadataInsight
	}{
		"EM_MANAGED_EXTERNAL_EXADATA",
		(MarshalTypeEmManagedExternalExadataInsight)(m),
	}

	return json.Marshal(&s)
}
